package log

import (
	"github.com/CameronHonis/set"
	"strings"
)

type Config struct {
	DecoratorByEnv map[string]func(string) string
	MutedEnvs      set.Set[string]
}

func NewConfig() *Config {
	return &Config{
		DecoratorByEnv: make(map[string]func(string) string),
		MutedEnvs:      *set.EmptySet[string](),
	}
}

func (c *Config) IsEnvDecorated(env string) bool {
	return c.DecoratorByEnv[strings.ToUpper(env)] != nil
}

func (c *Config) GetEnvDecorator(env string) func(string) string {
	return c.DecoratorByEnv[strings.ToUpper(env)]
}

func (c *Config) IsEnvMuted(env string) bool {
	return c.MutedEnvs.Has(strings.ToUpper(env))
}
