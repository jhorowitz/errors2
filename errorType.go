package errors2

import (
	"time"
	"fmt"
)

type LogMessage struct {
	time    time.Time
	message string
}

type FunctionTimer struct {
	FunctionName string
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
	err = FunctionTimingNeverStarted
	return
}

func (f *FunctionTimer) Report() (report string) {
	if len(f.FunctionName) == 0 {
		f.FunctionName = "anonymous"
	}
	if !f.started {
		return "FunctionTiming was never started."
	}
	if !f.ended {
		return fmt.Sprintf("FunctionTiming was never completed for %s. It's been %s since it started.", f.FunctionName, time.Since(f.startTime).String())
	}
	report += fmt.Sprintf("[%s]: completed in %s", f.FunctionName, f.endTime.Sub(f.startTime).String())
	return
}

func (f *FunctionTimer) Begin() {
	f.started = true
	f.startTime = time.Now()
}

func (f *FunctionTimer) Complete() {
	f.endTime = time.Now()
	f.ended = true
}

type Error struct {
	failureReason string
	failed        bool
}

func (e *Error) Fail(reason string) {
	e.failed = true
	e.failureReason = reason
}

func (e *Error) ToError() error {
	if !e.failed {
		return nil
	}
	return e
}

func (e *Error) Error() string {
	return e.failureReason
}

func NewFailedError(reason string) *Error {
	return &Error{failureReason: reason, failed:true}
}