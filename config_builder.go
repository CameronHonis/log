package log

import "strings"

type ConfigBuilder struct {
	Config *LoggerConfig
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		Config: NewLoggerConfig(false, make([]DecoratorRule, 0), make([]MutedRule, 0)),
	}
}

func (builder *ConfigBuilder) WithIsMuted(isMuted bool) *ConfigBuilder {
	builder.Config.IsMuted = isMuted
	return builder
}

func (builder *ConfigBuilder) WithDecorator(env string, decorator Decorator) *ConfigBuilder {
	rule := func(_env string) Decorator {
		if strings.ToUpper(_env) == strings.ToUpper(env) {
			return decorator
		}
		return nil
	}
	return builder.WithDecoratorRule(rule)
}

func (builder *ConfigBuilder) WithDecoratorRule(rule DecoratorRule) *ConfigBuilder {
	builder.Config.DecoratorRules = append(builder.Config.DecoratorRules, rule)
	return builder
}

func (builder *ConfigBuilder) WithMutedEnv(env string) *ConfigBuilder {
	rule := func(_env string) bool {
		return strings.ToUpper(env) == strings.ToUpper(_env)
	}
	return builder.WithMutedRule(rule)
}

func (builder *ConfigBuilder) WithMutedRule(rule MutedRule) *ConfigBuilder {
	builder.Config.MutedRules = append(builder.Config.MutedRules, rule)
	return builder
}

func (builder *ConfigBuilder) FromConfig(config *LoggerConfig) *ConfigBuilder {
	builder.Config = config
	return builder
}

func (builder *ConfigBuilder) WithConfig(config *LoggerConfig) *ConfigBuilder {
	for _, rule := range config.DecoratorRules {
		builder.WithDecoratorRule(rule)
	}
	for _, rule := range config.MutedRules {
		builder.WithMutedRule(rule)
	}
	return builder
}

func (builder *ConfigBuilder) Build() *LoggerConfig {
	return builder.Config
}
