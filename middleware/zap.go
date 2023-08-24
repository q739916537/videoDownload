package middleware

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func DefaultLog() *zap.Logger {
	if log == nil {
		fmt.Println("log is nil ,again init log middleware")
		InitLog()
	}
	return log
}

func InitLog() {
	// 配置 Zap 日志
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel), // 日志级别
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := cfg.Build()
	if err != nil {
		fmt.Println("Failed to create logger:", err)
		return
	}
	defer logger.Sync() // 刷新日志缓冲

	// 读取配置文件中的日志级别
	logLevel := cfg.Level.Level()
	switch logLevel {
	case zapcore.DebugLevel:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case zapcore.InfoLevel:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case zapcore.ErrorLevel:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		fmt.Println("Invalid log level in config")
		return
	}
	// 使用配置好的 Logger
	logger.Info("Logger initialized", zap.String("log_level", logLevel.String()))
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Error("This is an error message")
}
