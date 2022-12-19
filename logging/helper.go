package logging

import (
	"fmt"
)

func Infof(format string, keyVals ...interface{}) {
	logger.Log(InfoLevel, msgKey, fmt.Sprintf(format, keyVals...))
}

func Errorf(format string, keyVals ...interface{}) {
	logger.Log(ErrorLevel, msgKey, fmt.Sprintf(format, keyVals...))
}

func Warnf(format string, keyVals ...interface{}) {
	logger.Log(WarnLevel, msgKey, fmt.Sprintf(format, keyVals...))
}

func Debugf(format string, keyVals ...interface{}) {
	logger.Log(DebugLevel, msgKey, fmt.Sprintf(format, keyVals...))
}
