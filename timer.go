package errors2

import (
	"fmt"
	"time"
)

type Timer struct {
	startTime    time.Time
	endTime      time.Time
	started      bool
	ended        bool
}

func (f *Timer) Duration() (duration time.Duration, err error) {
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

func (f *Timer) Report() (report string) {
	if !f.started {
		return "Timing was never started."
	}
	if !f.ended {
		return "Timing was never completed."
	}
	report = fmt.Sprintf("completed in %s", f.endTime.Sub(f.startTime).String())
	return
}

func (f *Timer) Begin() {
	f.started = true
	f.startTime = time.Now()
}

func (f *Timer) Complete() {
	f.endTime = time.Now()
	f.ended = true
}
