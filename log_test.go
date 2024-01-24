package log

import (
	"bytes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"os"
)

func ReadStdout(stdoutWriter *os.File, stdoutReader *os.File) string {
	var stdoutBuffer bytes.Buffer
	_ = stdoutWriter.Close()
	_, _ = io.Copy(&stdoutBuffer, stdoutReader)
	return stdoutBuffer.String()
}

var _ = Describe("Log", func() {
	var log *Log
	var loggerConfig *LoggerConfig
	BeforeEach(func() {
		log = &Log{
			Env:          "TEST",
			Msg:          "test message",
			ColorWrapper: nil,
			Options:      make([]LogOption, 0),
		}
		loggerConfig = NewLoggerConfig(make([]DecoratorRule, 0), make([]MutedRule, 0))
	})
	Describe("::formatEnv", func() {
		It("formats the env", func() {
			Expect(log.formatEnv(loggerConfig)).To(Equal("[TEST]"))
		})
		When("the env is one of the designated envs", func() {
			BeforeEach(func() {
				loggerConfig = NewConfigBuilder().WithDecorator("server", WrapGreen).Build()

				log.Env = "SERVER"
			})
			It("formats the env with the designated color", func() {
				Expect(log.formatEnv(loggerConfig)).To(Equal("\x1b[32m[SERVER]\x1b[0m"))
			})
		})
	})
	Describe("::String", func() {
		It("returns the formatted log", func() {
			Expect(log.String(loggerConfig)).To(Equal("[TEST] test message"))
		})
		When("the log has a color wrapper", func() {
			It("applies the color wrapper to only the base message and returns the formatted log", func() {
				log.ColorWrapper = WrapRed
				Expect(log.String(loggerConfig)).To(Equal("[TEST] \x1b[31mtest message\x1b[0m"))
			})
		})
	})
})

var _ = Describe("LoggerService", func() {
	var oldStdout *os.File
	var stdoutWriter *os.File
	var stdoutReader *os.File
	var loggerService *LoggerService
	BeforeEach(func() {
		config := NewLoggerConfig(make([]DecoratorRule, 0), make([]MutedRule, 0))
		loggerService = NewLoggerService(config)
		oldStdout = os.Stdout
		stdoutReader, stdoutWriter, _ = os.Pipe()
		os.Stdout = stdoutWriter
	})
	Describe("::Log", func() {
		It("logs the message", func() {
			loggerService.Log("TEST", "test message")
			stdout := ReadStdout(stdoutWriter, stdoutReader)
			Expect(stdout).To(ContainSubstring("[TEST] test message\n"))
		})
		When("multiple strings are passed in", func() {
			It("logs the message", func() {
				loggerService.Log("TEST", "test message", " other test message", " 123")
				stdout := ReadStdout(stdoutWriter, stdoutReader)
				Expect(stdout).To(ContainSubstring("[TEST] test message other test message 123\n"))
			})
		})
		When("the env is muted", func() {
			BeforeEach(func() {
				config := NewConfigBuilder().WithMutedEnv("TEST").Build()
				loggerService = NewLoggerService(config)
			})
			It("does not log the message", func() {
				loggerService.Log("test", "test message")
				stdout := ReadStdout(stdoutWriter, stdoutReader)
				Expect(stdout).To(Equal(""))
			})
		})
	})
	Describe("::LogRed", func() {
		It("logs the message in color", func() {
			loggerService.LogRed("TEST", "test message")
			stdout := ReadStdout(stdoutWriter, stdoutReader)
			Expect(stdout).To(ContainSubstring("[TEST] \x1b[31mtest message\x1b[0m\n"))
		})
		When("the env is muted", func() {
			BeforeEach(func() {
				config := NewConfigBuilder().WithMutedEnv("TEST").Build()
				loggerService = NewLoggerService(config)
			})
			It("does not log the message", func() {
				loggerService.Log("TEST", "test message")
				stdout := ReadStdout(stdoutWriter, stdoutReader)
				Expect(stdout).To(Equal(""))
			})
		})
	})
	Describe("::LogGreen", func() {
		It("logs the message in color", func() {
			loggerService.LogGreen("TEST", "test message")
			stdout := ReadStdout(stdoutWriter, stdoutReader)
			Expect(stdout).To(ContainSubstring("[TEST] \x1b[32mtest message\x1b[0m\n"))
		})
	})
	AfterEach(func() {
		loggerService.LogBlue("TEST", "resetting stdout")
		_ = stdoutWriter.Close()
		_ = stdoutReader.Close()
		os.Stdout = oldStdout
	})
})
