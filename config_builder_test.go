package log_test

import (
	. "github.com/CameronHonis/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
			Expect(configBuilder.Config.DecoratorByEnv["TEST"]).ToNot(BeNil())
			Expect(configBuilder.Config.MutedEnvs.Has("MUTED_ENV")).To(BeTrue())
			configToMergeBuilder := NewConfigBuilder()
			configToMergeBuilder.WithDecorator("TEST2", WrapGreen)
			configToMergeBuilder.WithMutedEnv("MUTED_ENV2")
			configToMerge = configToMergeBuilder.Build()
		})
		It("contains the union of both decoratorByEnv map entries", func() {
			configBuilder.WithConfig(configToMerge)
			config := configBuilder.Build()
			Expect(len(config.DecoratorByEnv)).To(Equal(2))
			Expect(config.DecoratorByEnv["TEST"]).ToNot(BeNil())
			Expect(config.DecoratorByEnv["TEST2"]).ToNot(BeNil())
		})
		It("contains the union of both mutedEnvs sets", func() {
			configBuilder.WithConfig(configToMerge)
			config := configBuilder.Build()
			Expect(config.MutedEnvs.Has("MUTED_ENV")).To(BeTrue())
			Expect(config.MutedEnvs.Has("MUTED_ENV2")).To(BeTrue())
		})
		When("both configs specify a decorator for the same env", func() {
			BeforeEach(func() {
				configToMerge.DecoratorByEnv["TEST"] = WrapBlue
			})
			It("chooses the decorator from the passed in Config", func() {
				config := configBuilder.WithConfig(configToMerge).Build()
				Expect(config.DecoratorByEnv["TEST"]).ToNot(BeNil())
				wrappedEnv := config.DecoratorByEnv["TEST"]("test")
				Expect(wrappedEnv).To(Equal("\x1b[34mtest\x1b[0m"))
			})
		})
	})
})
