package log_type

type LogType string

var InfoType = InitLogType("info")
var WarnType = InitLogType("warn")
var ErrorType = InitLogType("error")
var FatalType = InitLogType("fatal")

func InitLogType(logType string) LogType {
	return LogType(logType)
}
