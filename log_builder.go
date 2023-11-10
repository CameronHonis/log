package log

import (
	"fmt"
	"strings"
)

type LogBuilder struct {
	log *Log
}

func NewLogBuilder() *LogBuilder {
	return &LogBuilder{
		log: &Log{},
	}
}

func (lb *LogBuilder) Env(env string) *LogBuilder {
	lb.log.Env = env
	return lb
}

func (lb *LogBuilder) Msg(msg string) *LogBuilder {
	lb.log.Msg = msg
	return lb
}

func (lb *LogBuilder) Msgs(msgs ...interface{}) *LogBuilder {
	var filteredMsg strings.Builder
	logOptions := make([]LogOption, 0, len(msgs))
	for _, msg := range msgs {
		if IsLogOption(msg) {
			logOptions = append(logOptions, LogOption(fmt.Sprintf("%v", msg)))
		} else {
			filteredMsg.WriteString(fmt.Sprintf("%v", msg))
		}
	}
	lb.log.Msg = filteredMsg.String()
	lb.log.Options = logOptions
	return lb
}

func (lb *LogBuilder) ColorWrapper(colorWrapper func(string) string) *LogBuilder {
	lb.log.ColorWrapper = colorWrapper
	return lb
}

func (lb *LogBuilder) Options(options ...LogOption) *LogBuilder {
	lb.log.Options = options
	return lb
}

func (lb *LogBuilder) Build() *Log {
	return lb.log
}
