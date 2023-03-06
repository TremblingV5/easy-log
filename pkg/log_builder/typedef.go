package log_builder

import "github.com/TremblingV5/easy-log/pkg/log_type"

type logInfo struct {
	key   string
	value string
}

type logInfoArray []*logInfo

type LogInfoBuilder struct {
	logType log_type.LogType
	message string
	Items   map[string]string
}
