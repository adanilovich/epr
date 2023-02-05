package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(logLevel, logPath string) (zap.Logger, error) {
	var zapLogLevel zapcore.Level
	switch logLevel {
	case "Debug":
		zapLogLevel = zapcore.DebugLevel
	case "Error":
		zapLogLevel = zapcore.ErrorLevel
	case "Info":
		zapLogLevel = zapcore.InfoLevel
	case "Warn":
		zapLogLevel = zapcore.WarnLevel
	default:
		zapLogLevel = zapcore.DebugLevel
	}

	c := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapLogLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    logLevel,
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
		OutputPaths: []string{logPath},
	}
	z, err := c.Build()
	if err != nil {
		return *z, err
	}

	return *z, nil
}
