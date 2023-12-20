package log

import (
	"github.com/CameronHonis/service"
	"github.com/CameronHonis/set"
	"strings"
)

type LoggerConfig struct {
	DecoratorByEnv map[string]func(string) string
	MutedEnvs      set.Set[string]
}

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		DecoratorByEnv: make(map[string]func(string) string),
		MutedEnvs:      *set.EmptySet[string](),
	}
}

func (c *LoggerConfig) MergeWith(mergingConfig service.ConfigI) service.ConfigI {
	_mergingConfig := mergingConfig.(*LoggerConfig)
	newConfig := *c
	for env, decorator := range _mergingConfig.DecoratorByEnv {
		newConfig.DecoratorByEnv[env] = decorator
	}
	for _, mutedEnv := range _mergingConfig.MutedEnvs.Flatten() {
		newConfig.MutedEnvs.Add(mutedEnv)
	}
	return &newConfig
}

func (c *LoggerConfig) IsEnvDecorated(env string) bool {
	return c.DecoratorByEnv[strings.ToUpper(env)] != nil
}

func (c *LoggerConfig) GetEnvDecorator(env string) func(string) string {
	return c.DecoratorByEnv[strings.ToUpper(env)]
}

func (c *LoggerConfig) IsEnvMuted(env string) bool {
	return c.MutedEnvs.Has(strings.ToUpper(env))
}
