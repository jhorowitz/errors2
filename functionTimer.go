package errors2

import (
	"fmt"
	"time"
)

type FunctionTimer struct {
	functionName string
	startTime    time.Time
	endTime      time.Time
	started      bool
	ended        bool
}

func (f *FunctionTimer) Duration() (duration time.Duration, err error) {
	if f.started && f.ended {
		duration = f.endTime.Sub(f.startTime)
		return
	}
	if !f.started {
		err = FunctionTimingNeverStarted
	} else if !f.ended {
		err = FunctionTimingNeverCompleted
	}
	return
}

func (f *FunctionTimer) Report() (report string) {
	if len(f.functionName) == 0 {
		f.functionName = "anonymous"
	}
	if !f.started {
		return "FunctionTiming was never started."
	}
	if !f.ended {
		return fmt.Sprintf("FunctionTiming was never completed for %s. It's been %s since it started.", f.functionName, time.Since(f.startTime).String())
	}
	report += fmt.Sprintf("[%s]: completed in %s", f.functionName, f.endTime.Sub(f.startTime).String())
	return
}

func (f *FunctionTimer) Begin(name string) {
	f.functionName = name
	f.started = true
	f.startTime = time.Now()
}

func (f *FunctionTimer) Complete() {
	f.endTime = time.Now()
	f.ended = true
}
