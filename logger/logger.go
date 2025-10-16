package logger

import "context"

type commonLog interface {
	Debugw(msg string, kv ...interface{})
	Infow(msg string, kv ...interface{})
	Warnw(msg string, kv ...interface{})
	Errorw(msg string, kv ...interface{})
	Fatalw(msg string, kv ...interface{})
	Debugf(msg string, kv ...interface{})
	Infof(msg string, kv ...interface{})
	Warnf(msg string, kv ...interface{})
	Errorf(msg string, kv ...interface{})
	Fatalf(msg string, kv ...interface{})
}

type Log interface {
	commonLog
	WithContext(ctx context.Context) Log
}
