/*
 * // Copyright 2020 Insolar Network Ltd.
 * // All rights reserved.
 * // This material is licensed under the Insolar License version 1.0,
 * // available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.
 */

package loaderbot

import (
	"context"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"go.uber.org/ratelimit"
)

const (
	DefaultMetricsUpdateInterval = 1 * time.Second
	DefaultScheduleQueueCapacity = 10000
	DefaultResultsQueueCapacity  = 10000
)

// Controlled struct for adding test vars
type Controlled struct {
	Sleep uint64
}

// TestData shared test data
type TestData struct {
	*sync.Mutex
	Index int
	Data  interface{}
}

type AsyncMetrics struct {
	Reported bool
	Samples  []AttackResult
	Metrics  *Metrics
}

type nextMsg struct {
	TargetRPS uint64
	Step      uint64
	Tick      uint64
}

// Runner provides test context for attacking target with constant amount of runners with a schedule
type Runner struct {
	// Name of a runner
	Name string
	// Cfg runner config
	Cfg *RunnerConfig
	// prototype from which all attackers cloned
	attackerPrototype Attack
	// target RPS for step, changed every step
	targetRPS int
	// current test step
	currentStep uint64
	// current metrics update tick
	currentTick uint64
	metricsMu   *sync.Mutex
	// metrics for every tick
	tickMetrics map[uint64]*AsyncMetrics

	rlMu *sync.Mutex
	rl   ratelimit.Limiter

	// TimeoutCtx test timeout ctx
	TimeoutCtx context.Context
	// test cancel func
	cancel context.CancelFunc
	// next schedule chan to signal to attack
	next chan nextMsg
	// attackers cloned for a prototype
	attackers []Attack

	results chan AttackResult
	// raw request results log
	rawResultsLog []AttackResult
	// uniq error messages
	uniqErrors map[string]int
	// Failed means there some errors in test
	Failed bool
	// data used to control attackers in test
	controlled Controlled
	// TestData data shared between attackers during test
	TestData interface{}
	L        *Logger
}

// NewRunner creates new runner with constant amount of attackers by RunnerConfig
func NewRunner(cfg *RunnerConfig, a Attack, data interface{}) *Runner {
	cfg.Validate()
	cfg.DefaultCfgValues()
	r := &Runner{
		Name:              cfg.Name,
		Cfg:               cfg,
		attackerPrototype: a,
		currentStep:       0,
		targetRPS:         cfg.StartRPS,
		metricsMu:         &sync.Mutex{},
		currentTick:       0,
		next:              make(chan nextMsg, DefaultScheduleQueueCapacity),
		rlMu:              &sync.Mutex{},
		rl:                ratelimit.New(cfg.StartRPS),
		attackers:         make([]Attack, 0),
		results:           make(chan AttackResult, DefaultResultsQueueCapacity),
		rawResultsLog:     make([]AttackResult, 0),
		tickMetrics:       make(map[uint64]*AsyncMetrics),
		uniqErrors:        make(map[string]int),
		controlled:        Controlled{},
		TestData:          data,
		L:                 NewLogger(cfg).With("runner", cfg.Name),
	}
	for i := 0; i < cfg.Attackers; i++ {
		a := r.attackerPrototype.Clone(r)
		if err := a.Setup(*r.Cfg); err != nil {
			log.Fatal(errAttackerSetup)
		}
		r.attackers = append(r.attackers, a)
	}
	return r
}

// Run runs the test
func (r *Runner) Run() (float64, error) {
	r.L.Infof("waiting for %d seconds before start", r.Cfg.WaitBeforeSec)
	time.Sleep(time.Duration(r.Cfg.WaitBeforeSec) * time.Second)
	r.L.Infof("runner started, mode: %s", r.Cfg.SystemMode)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.Cfg.TestTimeSec)*time.Second)
	r.TimeoutCtx = ctx
	r.cancel = cancel
	wg := &sync.WaitGroup{}
	wg.Add(len(r.attackers))
	for atkIdx, attacker := range r.attackers {
		r.L.Infof("starting attacker: %d", atkIdx)
		switch r.Cfg.SystemMode {
		case OpenWorldSystem:
			go asyncAttack(attacker, r, wg)
		case PrivateSystem:
			go attack(attacker, r, wg)
		}
	}
	wg.Wait()
	r.handleShutdownSignal()
	r.schedule()
	r.rampUp()
	r.tick()
	r.collectResults()
	<-r.TimeoutCtx.Done()
	r.cancel()
	r.L.Infof("runner exited")
	maxRPS := r.maxRPS()
	r.L.Infof("max rps: %.2f", maxRPS)
	return maxRPS, nil
}

func (r *Runner) printErrors() {
	r.L.Infof("Uniq errors:")
	for e, count := range r.uniqErrors {
		r.L.Infof("error: %s, count: %d", e, count)
	}
}

func (r *Runner) maxRPS() float64 {
	r.metricsMu.Lock()
	defer r.metricsMu.Unlock()
	rates := make([]float64, 0)
	for _, m := range r.tickMetrics {
		rates = append(rates, m.Metrics.Rate)
	}
	return MaxRPS(rates)
}

// schedule creates schedule plan for a test
func (r *Runner) schedule() {
	go func() {
		for {
			select {
			case <-r.TimeoutCtx.Done():
				r.L.Infof("schedule stopped")
				return
			default:
				r.rlMu.Lock()
				r.rl.Take()
				r.rlMu.Unlock()
				currentStep := atomic.LoadUint64(&r.currentStep)
				currentTick := atomic.LoadUint64(&r.currentTick)
				r.metricsMu.Lock()
				r.next <- nextMsg{
					TargetRPS: uint64(r.targetRPS),
					Step:      currentStep,
					Tick:      currentTick,
				}
				r.metricsMu.Unlock()
			}
		}
	}()
}

// rampUp changes ratelimit options on the run, increasing rate by StepRPS
func (r *Runner) rampUp() {
	ticker := time.NewTicker(time.Duration(r.Cfg.StepDurationSec) * time.Second)
	go func() {
		for {
			select {
			case <-r.TimeoutCtx.Done():
				return
			case <-ticker.C:
				currentStep := atomic.LoadUint64(&r.currentStep)
				atomic.AddUint64(&r.currentStep, 1)

				r.metricsMu.Lock()
				r.targetRPS += r.Cfg.StepRPS
				r.rlMu.Lock()
				r.rl = ratelimit.New(r.targetRPS)
				r.rlMu.Unlock()
				r.L.Infof("next step: step -> %d, rps -> %d", currentStep+1, r.targetRPS)
				r.metricsMu.Unlock()
				r.L.Infof("current active goroutines: %d", runtime.NumGoroutine())
			}
		}
	}()
}

// isCompleteTick tick is complete when 95% of samples received
func isCompleteTick(m *AsyncMetrics) bool {
	return float64(len(m.Samples)) >= float64(m.Samples[0].nextMsg.TargetRPS)*0.95 && !m.Reported
}

// collectResults collects attackers results and writes them to one of report options
func (r *Runner) collectResults() {
	go func() {
		for {
			select {
			case <-r.TimeoutCtx.Done():
				r.printErrors()
				return
			case res := <-r.results:
				if _, ok := r.tickMetrics[res.nextMsg.Tick]; !ok {
					r.tickMetrics[res.nextMsg.Tick] = &AsyncMetrics{
						false,
						make([]AttackResult, 0),
						NewMetrics(),
					}
				}
				currentTickMetrics := r.tickMetrics[res.nextMsg.Tick]
				currentTickMetrics.Samples = append(currentTickMetrics.Samples, res)
				if isCompleteTick(currentTickMetrics) {
					currentTickMetrics.Reported = true
					for _, s := range currentTickMetrics.Samples {
						currentTickMetrics.Metrics.add(s)
					}
					currentTickMetrics.Metrics.update()
					r.L.Infof("tick: %d, rate [%4f -> %v], perc: 50 [%v] 95 [%v], # requests [%d], %% success [%d]",
						res.nextMsg.Tick,
						currentTickMetrics.Metrics.Rate,
						res.nextMsg.TargetRPS,
						currentTickMetrics.Metrics.Latencies.P50,
						currentTickMetrics.Metrics.Latencies.P95,
						currentTickMetrics.Metrics.Requests,
						currentTickMetrics.Metrics.successLogEntry(),
					)
				}

				if res.doResult.Error != nil {
					r.uniqErrors[res.doResult.Error.Error()] += 1
					r.L.Debugf("attacker error: %s", res.doResult.Error)
					if r.Cfg.FailOnFirstError {
						r.Failed = true
						r.cancel()
					}
				}
				r.L.Debugf("received result: %v", res)
				r.rawResultsLog = append(r.rawResultsLog, res)
			}
		}
	}()
}

// tick emit ticks for metrics
func (r *Runner) tick() {
	ticker := time.NewTicker(DefaultMetricsUpdateInterval)
	go func() {
		for {
			select {
			case <-r.TimeoutCtx.Done():
				return
			case <-ticker.C:
				atomic.AddUint64(&r.currentTick, 1)
			}
		}
	}()
}
