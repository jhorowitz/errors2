package errors2

import (
	"testing"
	"time"
)

func TestTimerFunctionality(t *testing.T) {
	const waitDuration = time.Millisecond * 10
	t.Log("TestTimerFunctionality")
	ft := &FunctionTimer{}
	func() {
		ft.Begin("Blah")
		defer ft.Complete()
		time.Sleep(1 * waitDuration)
	}()
	duration, err := ft.Duration()
	if err != nil {
		t.Error(err.Error())
	}
	if duration < waitDuration {
		t.Error("Incorrect duration. Expected a time greater than or equal to " + time.Second.String() + " Got: " + duration.String())
	}
}
