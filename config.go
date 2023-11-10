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

func (lmc *LogManagerConfig) NormalizeConfig() {
	// upper case env names
	for env, decorator := range lmc.DecoratorByEnv {
		lmc.DecoratorByEnv[strings.ToUpper(env)] = decorator
	}
	for _, mutedEnv := range lmc.MutedEnvs.Flatten() {
		lmc.MutedEnvs.Remove(mutedEnv)
		lmc.MutedEnvs.Add(strings.ToUpper(mutedEnv))
	}
}
