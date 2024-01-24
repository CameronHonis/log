package log_test

import (
	. "github.com/CameronHonis/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("ConfigBuilder", func() {
	Describe("::FromConfig", func() {
		BeforeEach(func() {

		})
		It("overrides the Config to be the passed in Config", func() {

		})
	})
	Describe("::WithConfig", func() {
		var configBuilder *ConfigBuilder
		var configToMerge *LoggerConfig
		BeforeEach(func() {
			configBuilder = NewConfigBuilder()
			configBuilder.WithDecorator("test", WrapRed)
			configBuilder.WithMutedEnv("muted_env")
			Expect(configBuilder.Config).ToNot(BeNil())
			Expect(configBuilder.Config.IsEnvDecorated("test")).ToNot(BeNil())
			Expect(configBuilder.Config.IsEnvMuted("muted_env")).To(BeTrue())
			configToMergeBuilder := NewConfigBuilder()
			configToMergeBuilder.WithDecorator("TEST2", WrapGreen)
			configToMergeBuilder.WithMutedEnv("MUTED_ENV2")
			configToMerge = configToMergeBuilder.Build()
		})
		It("contains the union of both decoratorByEnv map entries", func() {
			configBuilder.WithConfig(configToMerge)
			config := configBuilder.Build()
			Expect(config.DecoratorRules).To(HaveLen(2))
			Expect(config.GetEnvDecorator("test")).ToNot(BeNil())
			Expect(config.GetEnvDecorator("test2")).ToNot(BeNil())
		})
		It("contains the union of both mutedEnvs sets", func() {
			configBuilder.WithConfig(configToMerge)
			config := configBuilder.Build()
			Expect(config.MutedRules).To(HaveLen(2))
			Expect(config.IsEnvMuted("muted_env")).To(BeTrue())
			Expect(config.IsEnvMuted("muted_env2")).To(BeTrue())
		})
		When("both configs specify a decorator for the same env", func() {
			BeforeEach(func() {
				decoratorRule := func(envName string) Decorator {
					if strings.ToUpper(envName) == "TEST" {
						return WrapBlue
					}
					return nil
				}
				configToMerge.DecoratorRules = append(configToMerge.DecoratorRules, decoratorRule)
			})
			It("chooses the decorator from the passed in Config", func() {
				config := configBuilder.WithConfig(configToMerge).Build()
				Expect(config.IsEnvDecorated("test")).ToNot(BeNil())
				wrappedEnv := config.GetEnvDecorator("test")("test")
				Expect(wrappedEnv).To(Equal("\x1b[34mtest\x1b[0m"))
			})
		})
	})
})
