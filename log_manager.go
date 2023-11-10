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
	lm.mu.Lock()
	lm.Config = *config
	lm.mu.Unlock()
}

func (lm *LogManager) logLogWithLock(log *Log) {
	lm.mu.Lock()
	fmt.Println(log.String())
	lm.mu.Unlock()
}

func (lm *LogManager) canPrintInEnv(msgs ...interface{}) bool {
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
	if !lm.canPrintInEnv(msgs...) {
		return
	}

	log := NewLogBuilder().Env(env).Msgs(msgs...).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogRed(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapRed).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogGreen(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapGreen).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogBlue(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapBlue).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogYellow(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapYellow).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogMagenta(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapMagenta).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogCyan(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapCyan).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogOrange(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapOrange).Build()
	lm.logLogWithLock(log)
}

func (lm *LogManager) LogBrown(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapBrown).Build()
	lm.logLogWithLock(log)
}