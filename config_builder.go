package log

import "strings"

type ConfigBuilder struct {
	Config *Config
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		Config: NewConfig(),
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

func (builder *ConfigBuilder) FromConfig(config *Config) *ConfigBuilder {
	builder.Config = config
	return builder
}

func (builder *ConfigBuilder) WithConfig(config *Config) *ConfigBuilder {
	for env, decorator := range config.DecoratorByEnv {
		builder.WithDecorator(env, decorator)
	}
	for _, mutedEnv := range config.MutedEnvs.Flatten() {
		builder.WithMutedEnv(mutedEnv)
	}
	return builder
}

func (builder *ConfigBuilder) Build() *Config {
	return builder.Config
}
