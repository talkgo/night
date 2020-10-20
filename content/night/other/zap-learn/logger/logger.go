package logger

import (
	"zap-learn/logger/zap"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})

	Fdebug(msg string, keyAndValues ...string)
	Finfo(msg string, keyAndValues ...string)
	Fwarn(msg string, keyAndValues ...string)
	Ferror(msg string, keyAndValues ...string)
	Fpanic(msg string, keyAndValues ...string)

	//DPanic(ars ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	//DPanicf(format string, ars ...interface{})

	TearDown()

	// 数据库引擎xorm的接口协议
	Level() int32
	SetLevel(l int32)
	ShowSQL(show ...bool)
	IsShowSQL() bool
}

var (
	// 这里需要导出Log变量是因为，数据库等第三方引擎需要自定制logger，可以直接使用该log
	Log Logger
)

func Setup() error {
	// Initialize global Log to ZLogger
	Log = Logger(zap.NewLogger())
	FInfo("logger setup succeed ^_^")
	return nil
}
func TearDown() {
	Log.TearDown()
}

func FDebug(prefix string, keyAndValues ...string) {
	Log.Fdebug(prefix, keyAndValues...)
	return
}

func FInfo(prefix string, keyAndValues ...string) {
	Log.Finfo(prefix, keyAndValues...)
	return
}

func FWarn(prefix string, keyAndValues ...string) {
	Log.Fwarn(prefix, keyAndValues...)
	return
}

func FError(prefix string, keyAndValues ...string) {
	Log.Ferror(prefix, keyAndValues...)
	return
}

func FPanic(prefix string, keyAndValues ...string) {
	Log.Fpanic(prefix, keyAndValues...)
	return
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	Log.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	Log.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	Log.Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message.
func Panic(args ...interface{}) {
	Log.Panic(args...)
}

// Debugf uses fmt.Sprint to construct and log a message.
func Debugf(fmt string, args ...interface{}) {
	Log.Debugf(fmt, args...)
}

// Infof uses fmt.Sprint to construct and log a message.
func Infof(fmt string, args ...interface{}) {
	Log.Infof(fmt, args...)
}

// Warnf uses fmt.Sprint to construct and log a message.
func Warnf(fmt string, args ...interface{}) {
	Log.Warnf(fmt, args...)
}

// Errorf uses fmt.Sprint to construct and log a message.
func Errorf(fmt string, args ...interface{}) {
	Log.Errorf(fmt, args...)
}

// Panicf uses fmt.Sprint to construct and log a message.
func Panicf(fmt string, args ...interface{}) {
	Log.Panicf(fmt, args...)
}
