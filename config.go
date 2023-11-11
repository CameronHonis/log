package log

import (
	"github.com/CameronHonis/set"
	"strings"
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

func (c *LogManagerConfig) IsEnvDecorated(env string) bool {
	return GetLogManager().Config.DecoratorByEnv[strings.ToUpper(env)] != nil
}

func (c *LogManagerConfig) GetEnvDecorator(env string) func(string) string {
	return GetLogManager().Config.DecoratorByEnv[strings.ToUpper(env)]
}

func (c *LogManagerConfig) IsEnvMuted(env string) bool {
	return GetLogManager().Config.MutedEnvs.Has(strings.ToUpper(env))
}
