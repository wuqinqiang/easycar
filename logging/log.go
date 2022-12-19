package logging

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

type Level int8

var msgKey = "msg"

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
)

type Logger interface {
	Log(level Level, keyvals ...interface{})
}

var logger Logger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)
	logger = &DefaultLog{logger: zap.New(core)}
}

// SetLog set custom logger
func SetLog(l Logger) {
	logger = l
}

// DefaultLog default logger by zap
type DefaultLog struct {
	logger *zap.Logger
}

func (d *DefaultLog) Log(level Level, keyvals ...interface{}) {
	if len(keyvals)%2 != 0 {
		d.logger.Warn(fmt.Sprintf("Keyvalues must appear in pairs:%v", keyvals))
		return
	}

	var (
		fields []zap.Field
	)

	if level != InfoLevel {
		fields = append(fields, getCallerInfoForLog()...)
	}

	for i := 0; i < len(keyvals); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}
	switch level {
	case InfoLevel:
		d.logger.Info("", fields...)
	case DebugLevel:
		d.logger.Debug("", fields...)
	case ErrorLevel:
		d.logger.Error("", fields...)
	case WarnLevel:
		d.logger.Warn("", fields...)
	default:
		d.logger.Warn(fmt.Sprintf("logging not included level:%v", level))
	}
}

func getCallerInfoForLog() (callerFields []zap.Field) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)

	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
	return
}
