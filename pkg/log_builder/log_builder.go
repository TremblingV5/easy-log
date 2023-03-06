package log_builder

import (
	"fmt"
	"github.com/TremblingV5/easy-log/pkg/log_type"
	"github.com/TremblingV5/easy-log/pkg/logger"
	"sort"
)

func InitLogBuilder(logType log_type.LogType, message string) *LogInfoBuilder {
	return &LogInfoBuilder{
		logType: logType,
		message: message,
		Items:   make(map[string]string),
	}
}

func (b *LogInfoBuilder) SetLogType(typeString log_type.LogType) {
	b.logType = typeString
}

func (b *LogInfoBuilder) SetMessage(str string) {
	b.message = str
}

func (b *LogInfoBuilder) Collect(key string, value string) {
	tempKey := key

	_, ok := b.Items[tempKey]
	if ok {
		cnt := 1
		for {
			_, ok := b.Items[tempKey+fmt.Sprint(cnt)]
			if !ok {
				tempKey = tempKey + fmt.Sprint(cnt)
				break
			}
			cnt++
		}
	}

	b.Items[tempKey] = value
}

func (b *LogInfoBuilder) Get() []string {
	temp := []*logInfo{}

	for k, v := range b.Items {
		newLogInfo := logInfo{
			key:   k,
			value: v,
		}
		temp = append(temp, &newLogInfo)
	}

	sort.Sort(logInfoArray(temp))
	res := []string{}
	for _, v := range temp {
		res = append(res, v.key)
		res = append(res, v.value)
	}
	return res
}

func (b *LogInfoBuilder) Write(logger *logger.Logger) {
	logger.Write(b.logType, b.message, b.Get()...)
}
