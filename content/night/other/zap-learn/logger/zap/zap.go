package zap

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZLogger struct {
	l  *zap.Logger
	sl *zap.SugaredLogger

	showSQL bool
}

// Debug uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Debug(args ...interface{}) {
	logger.sl.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Info(args ...interface{}) {
	logger.sl.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Warn(args ...interface{}) {
	logger.sl.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Error(args ...interface{}) {
	logger.sl.Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Panic(args ...interface{}) {
	if os.Getenv("RUN-MODE") == "development" {
		logger.sl.DPanic(args...)
	} else {
		logger.sl.Panic(args...)
	}
}
func (logger ZLogger) DPanic(template string, args ...interface{}) {
	return
}

func (logger ZLogger) Fdebug(msg string, keyAndValues ...string) {
	var fields []zap.Field
	if len(keyAndValues) > 0 {
		fields = parseFields(keyAndValues...)
	}
	logger.l.Debug(msg, fields...)
}

func (logger ZLogger) Finfo(msg string, keyAndValues ...string) {
	var fields []zap.Field
	if len(keyAndValues) > 0 {
		fields = parseFields(keyAndValues...)
	}
	logger.l.Info(msg, fields...)
}

func (logger ZLogger) Fwarn(msg string, keyAndValues ...string) {
	var fields []zap.Field
	if len(keyAndValues) > 0 {
		fields = parseFields(keyAndValues...)
	}
	logger.l.Warn(msg, fields...)
}

func (logger ZLogger) Ferror(msg string, keyAndValues ...string) {
	var fields []zap.Field
	if len(keyAndValues) > 0 {
		fields = parseFields(keyAndValues...)
	}
	logger.l.Error(msg, fields...)
}

func (logger ZLogger) Fpanic(msg string, keyAndValues ...string) {
	var fields []zap.Field
	if len(keyAndValues) > 0 {
		fields = parseFields(keyAndValues...)
	}
	if os.Getenv("RUN-MODE") == "development" {
		logger.l.DPanic(msg, fields...)
	} else {
		logger.l.Panic(msg, fields...)
	}
}

func parseFields(keyAndValues ...string) []zap.Field {
	var tempKey string
	fields := make([]zap.Field, 0, len(keyAndValues)/2)
	for index, keyOrValue := range keyAndValues {
		if index%2 == 1 {
			fields = append(fields, zap.String(tempKey, keyOrValue))
		} else {
			tempKey = keyOrValue
		}
	}
	return fields
}

// Debugf uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Debugf(template string, args ...interface{}) {
	logger.sl.Debugf(template, args...)
}

// Infof uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Infof(template string, args ...interface{}) {
	logger.sl.Infof(template, args...)
}

// Warnf uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Warnf(template string, args ...interface{}) {
	logger.sl.Warnf(template, args...)
}

// Errorf uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Errorf(template string, args ...interface{}) {
	logger.sl.Errorf(template, args...)
}

// Panicf uses fmt.Sprint to construct and log a message.
func (logger ZLogger) Panicf(template string, args ...interface{}) {
	if os.Getenv("RUN-MODE") == "development" {
		logger.sl.DPanicf(template, args...)
	} else {
		logger.sl.Panicf(template, args...)
	}
}

func (logger ZLogger) DPanicf(template string, args ...interface{}) {
	if os.Getenv("RUN-MODE") == "development" {
		logger.sl.DPanicf(template, args...)
	} else {
		logger.sl.Panicf(template, args...)
	}
}

func (logger ZLogger) TearDown() {
	logger.l.Sync()
}

func (logger ZLogger) Level() int32 {
	//if globalvar.DebugMode() {
	//	return core.LOG_DEBUG
	//} else {
	//	return core.LOG_INFO
	//}
	if os.Getenv("RUN-MODE") == "development" {
		return 1
	}
	return 0
}

// SetLevel 这个是xorm数据的接口实现，我们默认不设置，使失效，用系统统一设置的level
func (logger ZLogger) SetLevel(level int32) {
	return
}

func (logger ZLogger) ShowSQL(show ...bool) {
	if len(show) <= 0 {
		logger.showSQL = false
	} else {
		logger.showSQL = show[0]
	}
	return
}

func (logger ZLogger) IsShowSQL() bool {
	return logger.showSQL
}

func NewLogger() *ZLogger {
	logger := &ZLogger{}
	if os.Getenv("RUN-MODE") == "development" {
		//w := zapcore.AddSync(&lumberjack.Logger{
		//	Filename:   "./foo.log",
		//	MaxSize:    100, // megabytes
		//	MaxBackups: 4,
		//	MaxAge:     28, // days
		//	Compress:   false,
		//})
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel
		})
		encConfig := zap.NewDevelopmentEncoderConfig()
		encConfig.EncodeTime = timeEncoder
		encoder := zapcore.NewConsoleEncoder(encConfig)
		core := zapcore.NewTee(
			//zapcore.NewCore(encoder, w, zap.DebugLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lowPriority),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), highPriority),
		)
		tempLog := zap.New(core, zap.AddCallerSkip(2))

		logger.l = tempLog
		logger.sl = tempLog.Sugar()
		// 把标准库log也输出到 zaplog 中
		log.SetFlags(log.Lshortfile)
		log.SetOutput(&logWriter{logger.l})
		return logger
	}

	configJSON := fmt.Sprintf(`{
	 "level": "%s",
	 "development": %s,
	 "encoding": "%s",
	 "outputPaths": ["stdout", "%s"],
	 "errorOutputPaths": ["stderr", "%s"],
	 "initialFields": {"server": "%s"},
	 "encoderConfig": {
	   "messageKey": "message",
	   "levelKey": "level",
	   "levelEncoder": "lowercase"
	 }
	}`,
		//viper.GetString("log.logger_level"),
		//viper.GetString("log.development"),
		//viper.GetString("log.logger_encoding"),
		//viper.GetString("log.logger_normal_file"),
		//viper.GetString("log.logger_error_file"),
		//viper.GetString("service_name"),
		"debug", // debug, info, warn, error dpanic panic fatal
		"true", // development mode: true, false
		"json", // encoder: json
		"out.log", // normal out file
		"error.log", // error out file
		"server-name", // initial field name
	)

	var cfg zap.Config
	if err := json.Unmarshal([]byte(configJSON), &cfg); err != nil {
		//log.Fatal("Init zap logger failed with err:", err)
		panic(err)
	}
	cfg.Sampling = &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = timeEncoder
	logTemp, err := cfg.Build(zap.AddCallerSkip(2))
	if err != nil {
		//log.Fatal("Init zap logger failed with err:", err)
		panic(err)
	}
	logger.l = logTemp
	logger.sl = logTemp.Sugar()
	// 把标准库log也输出到 zaplog 中
	log.SetFlags(log.Lshortfile)
	log.SetOutput(&logWriter{logger.l})

	return logger
}

// callerEncoder will add caller to log. format is "filename:lineNum:funcName", e.g:"zaplog/zaplog_test.go:15:zaplog.TestNewLogger"
func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
}

// timeEncoder specifics the time format
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// milliSecondsDurationEncoder serializes a time.Duration to a floating-point number of micro seconds elapsed.
func milliSecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Millisecond))
}

func newLoggerConfig(debugLevel bool, te zapcore.TimeEncoder, de zapcore.DurationEncoder) (loggerConfig zap.Config) {
	loggerConfig = zap.NewProductionConfig()
	if te == nil {
		loggerConfig.EncoderConfig.EncodeTime = timeEncoder
	} else {
		loggerConfig.EncoderConfig.EncodeTime = te
	}
	if de == nil {
		loggerConfig.EncoderConfig.EncodeDuration = milliSecondsDurationEncoder
	} else {
		loggerConfig.EncoderConfig.EncodeDuration = de
	}
	loggerConfig.EncoderConfig.EncodeCaller = callerEncoder
	if debugLevel {
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	return
}

// NewLogger return a normal logger
func NewLoggerWithLevel(debugLevel bool) (logger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel, nil, nil)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewCustomLogger return a normal logger with given timeEncoder
func NewCustomLogger(debugLevel bool, te zapcore.TimeEncoder, de zapcore.DurationEncoder) (logger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel, te, de)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewNoCallerLogger return a no caller key value, will be faster
func NewNoCallerLogger(debugLevel bool) (noCallerLogger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel, nil, nil)
	loggerConfig.DisableCaller = true
	noCallerLogger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewNormalLoggers is a shortcut to get normal logger, noCallerLogger.
func NewNormalLoggers(debugLevel bool) (logger, noCallerLogger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel, nil, nil)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	loggerConfig.DisableCaller = true
	noCallerLogger, err = loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

type logWriter struct {
	logger *zap.Logger
}

// Write implement io.Writer, as std log's output
func (w logWriter) Write(p []byte) (n int, err error) {
	//i := bytes.Index(p, []byte(":")) + 1
	//j := bytes.Index(p[i:], []byte(":")) + 1 + i
	//caller := bytes.TrimRight(p[:j], ":")
	//// find last index of /
	//i = bytes.LastIndex(caller, []byte("/"))
	//// find penultimate index of /
	//i = bytes.LastIndex(caller[:i], []byte("/"))
	w.logger.Info("stdLog" + string(p))
	n = len(p)
	err = nil
	return
}
