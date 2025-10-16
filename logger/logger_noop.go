package logger

import (
	"context"
)

type NoopLogger struct{}

func NewNoop() NoopLogger                          { return NoopLogger{} }
func (NoopLogger) Debugw(string, ...interface{})   {}
func (NoopLogger) Infow(string, ...interface{})    {}
func (NoopLogger) Warnw(string, ...interface{})    {}
func (NoopLogger) Errorw(string, ...interface{})   {}
func (NoopLogger) Fatalw(string, ...interface{})   {}
func (NoopLogger) Debugf(string, ...interface{})   {}
func (NoopLogger) Infof(string, ...interface{})    {}
func (NoopLogger) Warnf(string, ...interface{})    {}
func (NoopLogger) Errorf(string, ...interface{})   {}
func (NoopLogger) Fatalf(string, ...interface{})   {}
func (NoopLogger) WithContext(context.Context) Log { return NoopLogger{} }
