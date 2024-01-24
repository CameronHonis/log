package log

import (
	"github.com/CameronHonis/service"
	"github.com/CameronHonis/set"
	"strings"
)

type DecoratorRule func(string) Decorator
type MutedRule func(string) bool

type LoggerConfig struct {
	service.ConfigI
	DecoratorRules []DecoratorRule
	MutedRules     []MutedRule
}

func NewLoggerConfig(decoratorRules []DecoratorRule, mutedRules []MutedRule) *LoggerConfig {
	return &LoggerConfig{
		DecoratorRules: decoratorRules,
		MutedRules:     mutedRules,
	}
}

func (c *LoggerConfig) IsEnvDecorated(env string) bool {
	return c.GetEnvDecorator(env) != nil
}

func (c *LoggerConfig) GetEnvDecorator(env string) Decorator {
	var decorator Decorator
	for _, rule := range c.DecoratorRules {
		if newDecorator := rule(env); newDecorator != nil {
			decorator = newDecorator
		}
	}
	return decorator
}

func (c *LoggerConfig) IsEnvMuted(env string) bool {
	for _, rule := range c.MutedRules {
		if rule(env) {
			return true
		}
	}
	return false
}

func NewDecoratorMapRule(decoratorByEnvName map[string]Decorator) DecoratorRule {
	decoratorByUpperEnvName := make(map[string]Decorator, len(decoratorByEnvName))
	for k, v := range decoratorByEnvName {
		decoratorByUpperEnvName[k] = v
	}

	rule := func(envName string) Decorator {
		return decoratorByUpperEnvName[strings.ToUpper(envName)]
	}
	return rule
}

func NewMutedSetRule(mutedSet *set.Set[string]) MutedRule {
	rule := func(envName string) bool {
		return mutedSet.Has(envName)
	}
	return rule
}
