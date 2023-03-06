package logger

import (
	"github.com/TremblingV5/easy-log/pkg/caller_skip"
	"github.com/TremblingV5/easy-log/pkg/log_type"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Handle *zap.Logger
}

func New(filename string, maxSize int, maxBackups int, maxAge int, skip caller_skip.CallerSkip) *Logger {
	writeSyncer := getLogWriter(
		filename, maxSize, maxBackups, maxAge,
	)

	encoder := getEncoder()

	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte("info")); err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(int(skip)))
	return &Logger{
		Handle: logger,
	}
}

func (l *Logger) Write(logType log_type.LogType, msg string, values ...string) {
	logInfo := make([]zapcore.Field, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		logInfo[i/2] = zap.String(values[i], values[i+1])
	}

	switch logType {
	case log_type.InfoType:
		l.Handle.Info(msg, logInfo...)
	case log_type.ErrorType:
		l.Handle.Error(msg, logInfo...)
	case log_type.WarnType:
		l.Handle.Warn(msg, logInfo...)
	case log_type.FatalType:
		l.Handle.Fatal(msg, logInfo...)
	default:
		l.Handle.Info(msg, logInfo...)
	}
}
