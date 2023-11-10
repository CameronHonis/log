package log

import "strings"

type LogManagerConfigBuilder struct {
	Config *LogManagerConfig
}

func NewLogManagerConfigBuilder() *LogManagerConfigBuilder {
	return &LogManagerConfigBuilder{
		Config: NewLogManagerConfig(),
	}
}

func (builder *LogManagerConfigBuilder) WithDecorator(env string, decorator func(string) string) *LogManagerConfigBuilder {
	env = strings.ToUpper(env)
	builder.Config.DecoratorByEnv[env] = decorator
	return builder
}

func (builder *LogManagerConfigBuilder) WithMutedEnv(env string) *LogManagerConfigBuilder {
	env = strings.ToUpper(env)
	builder.Config.MutedEnvs.Add(env)
	return builder
}

func (builder *LogManagerConfigBuilder) FromConfig(config *LogManagerConfig) *LogManagerConfigBuilder {
	builder.Config = config
	return builder
}

func (builder *LogManagerConfigBuilder) WithConfig(config *LogManagerConfig) *LogManagerConfigBuilder {
	for env, decorator := range config.DecoratorByEnv {
		builder.WithDecorator(env, decorator)
	}
	for _, mutedEnv := range config.MutedEnvs.Flatten() {
		builder.WithMutedEnv(mutedEnv)
	}
	return builder
}

func (builder *LogManagerConfigBuilder) Build() *LogManagerConfig {
	return builder.Config
}
