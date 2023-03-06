package main

import (
	"github.com/TremblingV5/easy-log/pkg/caller_skip"
	"github.com/TremblingV5/easy-log/pkg/log_builder"
	"github.com/TremblingV5/easy-log/pkg/log_type"
	"github.com/TremblingV5/easy-log/pkg/logger"
)

var logHandle = logger.New(
	"./temp.log", 1024*10, 3, 10, caller_skip.OnApp,
)

func solution() {
	logBuilder := log_builder.InitLogBuilder(
		log_type.InfoType, "Success",
	)
	defer logBuilder.Write(logHandle)

	logBuilder.SetMessage("Test log in a solution of an app")

	logBuilder.Collect("func name", "solution")

	if true {
		return
	}

	return
}

func main() {
	logBuilder := log_builder.InitLogBuilder(
		log_type.InfoType, "Success",
	)
	defer logBuilder.Write(logHandle)

	logBuilder.SetLogType(log_type.WarnType)
	logBuilder.SetMessage("Test log in warn type")

	logBuilder.Collect("timestamp", "1234567")
	logBuilder.Collect("name", "trembling")

	solution()
}
