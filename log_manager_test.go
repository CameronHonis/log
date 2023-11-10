package log_test

import (
	. "github.com/CameronHonis/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LogManager", func() {
	Describe("::InjectConfig", func() {
		var config *LogManagerConfig
		BeforeEach(func() {
			config = NewLogManagerConfigBuilder().WithDecorator("test", WrapRed).WithMutedEnv("test").Build()
		})
		JustBeforeEach(func() {
			GetLogManager().InjectConfig(config)
		})
		It("auto-capitalizes the env names", func() {
			Expect(GetLogManager().Config.DecoratorByEnv["TEST"]).ToNot(BeNil())
			Expect(GetLogManager().Config.MutedEnvs.Has("TEST")).To(BeTrue())
		})
		When("existing Config has been set", func() {
			var config2 *LogManagerConfig
			BeforeEach(func() {
				config2 = NewLogManagerConfig()
				config2.DecoratorByEnv["test2"] = WrapGreen
				config2.MutedEnvs.Add("test2")
			})
			JustBeforeEach(func() {
				GetLogManager().InjectConfig(config2)
			})
			It("adds the new Config", func() {
				Expect(GetLogManager().Config.DecoratorByEnv["TEST2"]).ToNot(BeNil())
				Expect(GetLogManager().Config.MutedEnvs.Has("TEST2")).To(BeTrue())
			})
			It("preserves existing Config", func() {
				Expect(GetLogManager().Config.DecoratorByEnv["TEST"]).ToNot(BeNil())
				Expect(GetLogManager().Config.MutedEnvs.Has("TEST")).To(BeTrue())
			})
		})
	})
	Describe("::WipeConfig", func() {
		BeforeEach(func() {
			config := NewLogManagerConfig()
			config.DecoratorByEnv["TEST"] = WrapRed
			config.MutedEnvs.Add("TEST")
			GetLogManager().InjectConfig(config)
			Expect(GetLogManager().Config.DecoratorByEnv["TEST"]).ToNot(BeNil())
			Expect(GetLogManager().Config.MutedEnvs.Has("TEST")).To(BeTrue())
		})
		JustBeforeEach(func() {
			GetLogManager().WipeConfig()
		})
		It("wipes the env decorators", func() {
			Expect(GetLogManager().Config.DecoratorByEnv["TEST"]).To(BeNil())
		})
		It("wipes the muted envs", func() {
			Expect(GetLogManager().Config.MutedEnvs.Has("TEST")).To(BeFalse())
		})
	})
})
