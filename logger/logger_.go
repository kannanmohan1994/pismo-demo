package logger

import (
	"context"
	"xm/config"
	"xm/consts"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap *zap.SugaredLogger
}

func NewLogger(cfg *config.Config) Log {
	logger, err := getLoggerbyEnv(cfg.Environment)
	if err != nil {
		panic(err)
	}

	return &Logger{logger.Sugar()}
}

func getLoggerbyEnv(env string) (logger *zap.Logger, err error) {
	option := zap.AddCallerSkip(1)

	if env == string(consts.PRODUCTION) {
		return zap.NewProduction(option)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return config.Build(option)
}

func (l *Logger) Debugw(msg string, kv ...interface{}) {
	l.zap.Debugw(msg, kv...)
}

func (l *Logger) Errorw(msg string, kv ...interface{}) {
	l.zap.Errorw(msg, kv...)
}

func (l *Logger) Fatalw(msg string, kv ...interface{}) {
	l.zap.Fatalw(msg, kv...)
}

func (l *Logger) Infow(msg string, kv ...interface{}) {
	l.zap.Infow(msg, kv...)
}

func (l *Logger) Warnw(msg string, kv ...interface{}) {
	l.zap.Warnw(msg, kv...)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.zap.Debugf(msg, args...)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.zap.Errorf(msg, args...)
}

func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.zap.Fatalf(msg, args...)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.zap.Infof(msg, args...)
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.zap.Warnf(msg, args...)
}

func (l *Logger) WithContext(ctx context.Context) Log {
	var newLogger *Logger

	if ctx != nil {
		if ctxRqId, ok := ctx.Value(consts.CorrelationID).(string); ok {
			newLogger = &Logger{
				zap: l.zap.With(zap.String(string(consts.CorrelationID), ctxRqId)),
			}
		}
	}

	return newLogger
}
