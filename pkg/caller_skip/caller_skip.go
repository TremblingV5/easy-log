package caller_skip

type CallerSkip int

// OnApp Make "caller" in formatted logs locate at the code in app
var OnApp = InitCallerSkip(2)

func InitCallerSkip(skip int) CallerSkip {
	return CallerSkip(skip)
}
