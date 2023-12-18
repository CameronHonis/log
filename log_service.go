package log

import (
	"fmt"
	"strings"
	"sync"
)

type LogManagerI interface {
	InjectConfig(config *Config)
	WipeConfig()
	Log(env string, msgs ...interface{})
	LogRed(env string, msgs ...interface{})
	LogGreen(env string, msgs ...interface{})
	LogBlue(env string, msgs ...interface{})
	LogYellow(env string, msgs ...interface{})
	LogMagenta(env string, msgs ...interface{})
	LogCyan(env string, msgs ...interface{})
	LogOrange(env string, msgs ...interface{})
	LogBrown(env string, msgs ...interface{})
}

var logService *LogService

type LogService struct {
	mu     sync.Mutex
	Config Config
}

func GetLogService() *LogService {
	if logService == nil {
		logService = &LogService{
			mu:     sync.Mutex{},
			Config: *NewConfig(),
		}
	}
	return logService
}

func (lm *LogService) InjectConfig(config *Config) {
	configBuilder := NewConfigBuilder().FromConfig(&lm.Config).WithConfig(config)
	lm.mu.Lock()
	lm.Config = *configBuilder.Build()
	lm.mu.Unlock()
}

func (lm *LogService) WipeConfig() {
	lm.mu.Lock()
	lm.Config = *NewConfig()
	lm.mu.Unlock()
}

func (lm *LogService) logLogWithLock(log *Log) {
	lm.mu.Lock()
	fmt.Println(log.String())
	lm.mu.Unlock()
}

func (lm *LogService) canPrintInEnv(env string, msgs ...interface{}) bool {
	// NOTE: env means two things here:
	// 1. the environment in which the log was created
	// 2. the runtime environment (i.e. test, prod, etc.)
	// TODO: rename one of these envs to avoid confusion

	if lm.Config.MutedEnvs.Has(strings.ToUpper(env)) {
		return false
	}
	for _, msg := range msgs {
		switch msg {
		case ONLY_TEST_ENV:
			return GetEnvName() == "test"
		case ONLY_PROD_ENV:
			return GetEnvName() == "prod"
		case ALL_BUT_TEST_ENV:
			return GetEnvName() != "test"
		}
	}
	return true
}

func (lm *LogService) Log(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}

	log := NewLogBuilder().Env(env).Msgs(msgs...).Build()
	lm.logLogWithLock(log)
}

func (lm *LogService) LogRed(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapRed).Build()
	lm.logLogWithLock(log)
}

func (lm *LogService) LogGreen(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapGreen).Build()
	lm.logLogWithLock(log)
}

func (lm *LogService) LogBlue(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapBlue).Build()
	lm.logLogWithLock(log)
}

func (lm *LogService) LogYellow(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapYellow).Build()
	lm.logLogWithLock(log)
}

func (lm *LogService) LogMagenta(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapMagenta).Build()
	lm.logLogWithLock(log)
}

func (lm *LogService) LogCyan(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapCyan).Build()
	lm.logLogWithLock(log)
}

func (lm *LogService) LogOrange(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapOrange).Build()
	lm.logLogWithLock(log)
}

func (lm *LogService) LogBrown(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapBrown).Build()
	lm.logLogWithLock(log)
}
