package logging

import (
	"fmt"
	"path"
	"runtime"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func Infof(format string, keyVals ...interface{}) {
	log(zapcore.InfoLevel, fmt.Sprintf(format, keyVals...))
}

func Errorf(format string, keyVals ...interface{}) {
	log(zapcore.ErrorLevel, fmt.Sprintf(format, keyVals...))
}

func Warnf(format string, keyVals ...interface{}) {
	log(zapcore.WarnLevel, fmt.Sprintf(format, keyVals...))
}

func Debugf(format string, keyVals ...interface{}) {
	log(zapcore.DebugLevel, fmt.Sprintf(format, keyVals...))
}

func log(level zapcore.Level, msg string, keyvals ...interface{}) {
	if len(keyvals)%2 != 0 {
		logger.Warn(fmt.Sprintf("Keyvalues must appear in pairs:%v", keyvals))
		return
	}

	var (
		fields []zap.Field
	)

	if level != zapcore.InfoLevel {
		fields = append(fields, getCallerInfoForLog()...)
	}

	for i := 0; i < len(keyvals); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}
	switch level {
	case zapcore.InfoLevel:
		logger.Info(msg, fields...)
	case zapcore.DebugLevel:
		logger.Debug(msg, fields...)
	case zapcore.FatalLevel:
		logger.Fatal(msg, fields...)
	case zapcore.ErrorLevel:
		logger.Error(msg, fields...)
	case zapcore.WarnLevel:
		logger.Warn(msg, fields...)
	default:
		logger.Warn(fmt.Sprintf("logging not included level:%v", level))
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
