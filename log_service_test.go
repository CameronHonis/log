package log_test

import (
	. "github.com/CameronHonis/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LogService", func() {
	Describe("::InjectConfig", func() {
		var config *Config
		BeforeEach(func() {
			config = NewConfigBuilder().WithDecorator("test", WrapRed).WithMutedEnv("test").Build()
		})
		JustBeforeEach(func() {
			GetLogService().InjectConfig(config)
		})
		It("auto-capitalizes the env names", func() {
			Expect(GetLogService().Config.DecoratorByEnv["TEST"]).ToNot(BeNil())
			Expect(GetLogService().Config.MutedEnvs.Has("TEST")).To(BeTrue())
		})
		When("existing Config has been set", func() {
			var config2 *Config
			BeforeEach(func() {
				config2 = NewConfig()
				config2.DecoratorByEnv["test2"] = WrapGreen
				config2.MutedEnvs.Add("test2")
			})
			JustBeforeEach(func() {
				GetLogService().InjectConfig(config2)
			})
			It("adds the new Config", func() {
				Expect(GetLogService().Config.DecoratorByEnv["TEST2"]).ToNot(BeNil())
				Expect(GetLogService().Config.MutedEnvs.Has("TEST2")).To(BeTrue())
			})
			It("preserves existing Config", func() {
				Expect(GetLogService().Config.DecoratorByEnv["TEST"]).ToNot(BeNil())
				Expect(GetLogService().Config.MutedEnvs.Has("TEST")).To(BeTrue())
			})
		})
	})
	Describe("::WipeConfig", func() {
		BeforeEach(func() {
			config := NewConfig()
			config.DecoratorByEnv["TEST"] = WrapRed
			config.MutedEnvs.Add("TEST")
			GetLogService().InjectConfig(config)
			Expect(GetLogService().Config.DecoratorByEnv["TEST"]).ToNot(BeNil())
			Expect(GetLogService().Config.MutedEnvs.Has("TEST")).To(BeTrue())
		})
		JustBeforeEach(func() {
			GetLogService().WipeConfig()
		})
		It("wipes the env decorators", func() {
			Expect(GetLogService().Config.DecoratorByEnv["TEST"]).To(BeNil())
		})
		It("wipes the muted envs", func() {
			Expect(GetLogService().Config.MutedEnvs.Has("TEST")).To(BeFalse())
		})
	})
})
