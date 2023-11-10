package log

import (
	"fmt"
	"sync"
)

var logManager *LogManager

type LogManager struct {
	mu     sync.Mutex
	Config LogManagerConfig
}

func GetLogManager() *LogManager {
	if logManager == nil {
		logManager = &LogManager{
			mu:     sync.Mutex{},
			Config: *NewLogManagerConfig(),
		}
	}
	return logManager
}

func (lm *LogManager) InjectConfig(config *LogManagerConfig) {
	configBuilder := NewLogManagerConfigBuilder().FromConfig(&lm.Config).WithConfig(config)
	lm.mu.Lock()
	lm.Config = *configBuilder.Build()
	lm.mu.Unlock()
}

func (lm *LogManager) WipeConfig() {
	lm.mu.Lock()
	lm.Config = *NewLogManagerConfig()
	lm.mu.Unlock()
}

func (lm *LogManager) logLogWithLock(log *Log) {
	lm.mu.Lock()
	fmt.Println(log.String())
	lm.mu.Unlock()
}

func (lm *LogManager) canPrintInEnv(env string, msgs ...interface{}) bool {
	// NOTE: env means two things here:
	// 1. the environment in which the log was created
	// 2. the runtime environment (i.e. test, prod, etc.)
	// TODO: rename one of these envs to avoid confusion

	if lm.Config.MutedEnvs.Has(env) {
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

func (lm *LogManager) Log(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}

	log := NewLogBuilder().Env(env).Msgs(msgs...).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogRed(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapRed).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogGreen(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapGreen).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogBlue(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapBlue).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogYellow(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapYellow).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogMagenta(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapMagenta).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogCyan(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapCyan).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogOrange(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapOrange).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogBrown(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapBrown).Build()
	lm.logLogWithLock(log)
}
