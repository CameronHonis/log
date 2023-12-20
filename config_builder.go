package log

import "strings"

type ConfigBuilder struct {
	Config *LoggerConfig
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		Config: NewLoggerConfig(),
	}
}

func (builder *ConfigBuilder) WithDecorator(env string, decorator func(string) string) *ConfigBuilder {
	env = strings.ToUpper(env)
	builder.Config.DecoratorByEnv[env] = decorator
	return builder
}

func (builder *ConfigBuilder) WithMutedEnv(env string) *ConfigBuilder {
	env = strings.ToUpper(env)
	builder.Config.MutedEnvs.Add(env)
	return builder
}

func (builder *ConfigBuilder) FromConfig(config *LoggerConfig) *ConfigBuilder {
	builder.Config = config
	return builder
}

func (builder *ConfigBuilder) WithConfig(config *LoggerConfig) *ConfigBuilder {
	for env, decorator := range config.DecoratorByEnv {
		builder.WithDecorator(env, decorator)
	}
	for _, mutedEnv := range config.MutedEnvs.Flatten() {
		builder.WithMutedEnv(mutedEnv)
	}
	return builder
}

func (builder *ConfigBuilder) Build() *LoggerConfig {
	return builder.Config
}
