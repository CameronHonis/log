package log

import (
	"fmt"
	"github.com/CameronHonis/marker"
	"github.com/CameronHonis/service"
	"sync"
)

type LoggerServiceI interface {
	service.ServiceI
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

type LoggerService struct {
	service.Service

	__dependencies__ marker.Marker

	__state__ marker.Marker
	mu        sync.Mutex
}

func NewLoggerService(config *LoggerConfig) *LoggerService {
	loggerService := &LoggerService{
		mu: sync.Mutex{},
	}
	loggerService.Service = *service.NewService(loggerService, config)
	return loggerService
}

func (lm *LoggerService) logLogWithLock(log *Log, config *LoggerConfig) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	if config.IsMuted {
		return
	}
	fmt.Println(log.String(config))
}

func (lm *LoggerService) canPrintInEnv(env string, msgs ...interface{}) bool {
	// NOTE: env refers to the interprocess environment, not anything related to the host machine environment.
	// TODO: consider renaming env to something better

	if lm.Config().(*LoggerConfig).IsEnvMuted(env) {
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

func (lm *LoggerService) Log(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}

	log := NewLogBuilder().Env(env).Msgs(msgs...).Build()
	lm.logLogWithLock(log, lm.Config().(*LoggerConfig))
}

func (lm *LoggerService) LogRed(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapRed).Build()
	lm.logLogWithLock(log, lm.Config().(*LoggerConfig))
}

func (lm *LoggerService) LogGreen(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapGreen).Build()
	lm.logLogWithLock(log, lm.Config().(*LoggerConfig))
}

func (lm *LoggerService) LogBlue(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapBlue).Build()
	lm.logLogWithLock(log, lm.Config().(*LoggerConfig))
}

func (lm *LoggerService) LogYellow(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapYellow).Build()
	lm.logLogWithLock(log, lm.Config().(*LoggerConfig))
}

func (lm *LoggerService) LogMagenta(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapMagenta).Build()
	lm.logLogWithLock(log, lm.Config().(*LoggerConfig))
}

func (lm *LoggerService) LogCyan(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapCyan).Build()
	lm.logLogWithLock(log, lm.Config().(*LoggerConfig))
}

func (lm *LoggerService) LogOrange(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapOrange).Build()
	lm.logLogWithLock(log, lm.Config().(*LoggerConfig))
}

func (lm *LoggerService) LogBrown(env string, msgs ...interface{}) {
	if !lm.canPrintInEnv(env, msgs...) {
		return
	}
	log := NewLogBuilder().Env(env).Msgs(msgs...).ColorWrapper(WrapBrown).Build()
	lm.logLogWithLock(log, lm.Config().(*LoggerConfig))
}
