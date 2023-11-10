package log

import (
	"github.com/CameronHonis/set"
)

type LogManagerConfig struct {
	DecoratorByEnv map[string]func(string) string
	MutedEnvs      set.Set[string]
}

func NewLogManagerConfig() *LogManagerConfig {
	return &LogManagerConfig{
		DecoratorByEnv: make(map[string]func(string) string),
		MutedEnvs:      *set.EmptySet[string](),
	}
}
