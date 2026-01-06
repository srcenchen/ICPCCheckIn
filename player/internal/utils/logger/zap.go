package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	once   sync.Once
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func Init() error {
	var err error
	once.Do(func() {
		config := zap.NewDevelopmentConfig()
		config.DisableStacktrace = true // 关闭堆栈
		logger, err = config.Build()
		if err == nil {
			sugar = logger.Sugar()
		}
	})
	return err
}

func Logger() *zap.Logger {
	if logger == nil {
		return zap.NewNop()
	}
	return logger
}

func Sugar() *zap.SugaredLogger {
	if logger == nil {
		return zap.NewNop().Sugar()
	}
	return sugar
}
