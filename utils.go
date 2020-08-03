/*
 * // Copyright 2020 Insolar Network Ltd.
 * // All rights reserved.
 * // This material is licensed under the Insolar License version 1.0,
 * // available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.
 */

package loaderbot

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

var (
	sigs = make(chan os.Signal, 1)
)

func (r *Runner) handleShutdownSignal() {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-r.TimeoutCtx.Done():
			return
		case <-sigs:
			if r.Cfg.ReportOptions.CSV {
				r.flushLogs()
			}
			r.cancel()
			r.L.Infof("exit signal received, exiting")
			if r.Cfg.GoroutinesDump {
				buf := make([]byte, 1<<20)
				stacklen := runtime.Stack(buf, true)
				r.L.Infof("=== received SIGTERM ===\n*** goroutine dump...\n%s\n*** End\n", buf[:stacklen])
			}
			os.Exit(1)
		}
	}()
}

// CreateFileOrAppend creates file if not exists or opens in append mode, used for metrics between tests
func CreateFileOrAppend(fname string) *os.File {
	var file *os.File
	fpath, _ := filepath.Abs(fname)
	_, err := os.Stat(fpath)
	if err != nil {
		file, err = os.Create(fname)
	} else {
		file, err = os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	if err != nil {
		log.Fatal(err)
	}
	return file
}

// CreateFileOrReplace creates new file every time, used for files with static name
// content of which must not contain data from different tests
func CreateFileOrReplace(fname string) *os.File {
	fpath, _ := filepath.Abs(fname)
	_ = os.Remove(fpath)
	file, err := os.Create(fpath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func MaxRPS(array []float64) float64 {
	if len(array) == 0 {
		return 1
	}
	var max = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}
