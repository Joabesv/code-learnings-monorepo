package logging

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Setup() {
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		panic("failed to create logs directory: " + err.Error())
	}

	logLevel := getLogLevel()

	jsonEncoderConfig := zap.NewProductionEncoderConfig()
	jsonEncoderConfig.TimeKey = "time"
	jsonEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	jsonEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	jsonEncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	jsonEncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(jsonEncoderConfig),
		zapcore.AddSync(getLogFile(logsDir, "app.log")),
		logLevel,
	)

	var consoleCore zapcore.Core
	if os.Getenv("APP_ENV") == "development" {
		consoleCore = zapcore.NewCore(
			zapcore.NewJSONEncoder(jsonEncoderConfig),
			zapcore.AddSync(os.Stdout),
			logLevel,
		)
	}

	var core zapcore.Core
	if consoleCore != nil {
		core = zapcore.NewTee(fileCore, consoleCore)
	} else {
		core = fileCore
	}

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)

	logger.Info("Logger initialized",
		zap.String("level", logLevel.String()),
		zap.String("environment", getEnvironment()),
		zap.String("log_file", filepath.Join(logsDir, "app.log")),
	)
}

func getLogLevel() zapcore.Level {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		return zapcore.InfoLevel
	}

	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		return zapcore.InfoLevel
	}
	return level
}

func getEnvironment() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return "production"
	}
	return env
}

func getLogFile(logsDir, filename string) *os.File {
	logPath := filepath.Join(logsDir, filename)

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}

	return file
}

func Sync() {
	if logger != nil {
		logger.Sync()
	}
}

func Info(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Info(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Error(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Warn(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Debug(msg, fields...)
	}
}

func Fatal(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Fatal(msg, fields...)
	}
}

func WithFields(fields ...zap.Field) *zap.Logger {
	if logger != nil {
		return logger.With(fields...)
	}
	return nil
}

func WithError(err error) *zap.Logger {
	if logger != nil {
		return logger.With(zap.Error(err))
	}
	return nil
}

func GetLogger() *zap.Logger {
	return logger
}

func L() *zap.Logger {
	return zap.L()
}
