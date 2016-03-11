package errors2

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