package log

import (
	"fmt"
	"os"
	"strings"
)

type Decorator func(string) string

type Log struct {
	Env          string
	Msg          string
	ColorWrapper Decorator
	Options      []LogOption
}

func (l *Log) formatEnv(config *LoggerConfig) string {
	envWrapper := config.GetEnvDecorator(l.Env)
	if envWrapper != nil {
		return envWrapper(fmt.Sprintf("[%s]", strings.ToUpper(l.Env)))
	}
	return fmt.Sprintf("[%s]", strings.ToUpper(l.Env))
}

func (l *Log) String(config *LoggerConfig) string {
	if l.ColorWrapper == nil {
		return fmt.Sprintf("%s %s", l.formatEnv(config), l.Msg)
	} else {
		return fmt.Sprintf("%s %s", l.formatEnv(config), l.ColorWrapper(l.Msg))
	}
}

type LogOption string

const (
	ONLY_TEST_ENV    LogOption = "ONLY_TEST_ENV"
	ONLY_PROD_ENV              = "ONLY_PROD_ENV"
	ALL_BUT_TEST_ENV           = "ALL_BUT_TEST_ENV"
)

func IsLogOption(m interface{}) bool {
	return m == ONLY_TEST_ENV || m == ONLY_PROD_ENV || m == ALL_BUT_TEST_ENV
}

func GetEnvName() string {
	envName, hasEnvVar := os.LookupEnv("ENV")
	if !hasEnvVar {
		envName = "test"
	}
	return strings.ToLower(envName)
}
