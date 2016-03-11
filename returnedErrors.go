package errors2

const prepend = "errors2: "

var FunctionTimingNeverStarted = (&Error{failureReason: prepend + "The timer was never started.", failed:true}).ToError()
var FunctionTimingNeverCompleted = (&Error{failureReason: prepend + "The timer is still running..", failed:true}).ToError()

