package errors2


func newFailedInternalError(reason string) *Error {
	const prepend = "errors2: "
	return NewFailedError(prepend + reason)
}

var (
	FunctionTimingNeverStarted = newFailedInternalError("The timer was never started.").ToError()
	FunctionTimingNeverCompleted = newFailedInternalError("The timer is still running..").ToError()
)