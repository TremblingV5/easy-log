package caller_skip

type CallerSkip int

var OnApp = InitCallerSkip(2)

func InitCallerSkip(skip int) CallerSkip {
	return CallerSkip(skip)
}
