package log_test

import (
	"github.com/CameronHonis/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("Config", func() {
	Describe("Decorators", func() {
		Describe("::GetEnvDecorator", func() {
			var config *log.LoggerConfig
			BeforeEach(func() {
				rule0Map := map[string]log.Decorator{
					"SOME_ENV_ONE": log.WrapRed,
					"SOME_ENV_TWO": log.WrapBlue,
				}
				rule0 := log.NewDecoratorMapRule(rule0Map)
				rule1 := func(envName string) log.Decorator {
					if envName == "SOME_ENV_TWO" {
						return log.WrapGreen
					}
					return nil
				}
				rules := []log.DecoratorRule{rule0, rule1}
				config = log.NewLoggerConfig(false, rules, make([]log.MutedRule, 0))
			})
			When("no rules specify the decorator for the env", func() {
				It("returns nil", func() {
					Expect(config.GetEnvDecorator("SOME_ENV_ZERO")).To(BeNil())
				})
			})
			When("one rule specifies the decorator for the env", func() {
				It("returns that decorator", func() {
					decoratorAddr := reflect.ValueOf(config.GetEnvDecorator("SOME_ENV_ONE")).Pointer()
					expAddr := reflect.ValueOf(log.WrapRed).Pointer()
					Expect(decoratorAddr).To(Equal(expAddr))
				})
			})
			When("two rules specify the decorator for the env", func() {
				It("returns the last decorator", func() {
					decoratorAddr := reflect.ValueOf(config.GetEnvDecorator("SOME_ENV_TWO")).Pointer()
					expAddr := reflect.ValueOf(log.WrapGreen).Pointer()
					Expect(decoratorAddr).To(Equal(expAddr))
				})
			})
		})
	})
})
