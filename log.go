package log

import (
	"fmt"
	"os"
	"strings"
)

type Log struct {
	Env          string
	Msg          string
	ColorWrapper func(string) string
	Options      []LogOption
}

func (l *Log) formatEnv() string {
	upperEnv := strings.ToUpper(l.Env)
	envWrapper, ok := GetLogService().Config.DecoratorByEnv[upperEnv]
	if ok {
		return envWrapper(fmt.Sprintf("[%s]", upperEnv))
	}
	return fmt.Sprintf("[%s]", upperEnv)
}

func (l *Log) String() string {
	if l.ColorWrapper == nil {
		return fmt.Sprintf("%s %s", l.formatEnv(), l.Msg)
	} else {
		return fmt.Sprintf("%s %s", l.formatEnv(), l.ColorWrapper(l.Msg))
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
