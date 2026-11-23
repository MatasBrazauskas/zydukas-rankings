package utils

import "go.uber.org/zap"

var zapLog *zap.SugaredLogger

func InitLogger() error {
	logger, sErr := zap.NewDevelopment()
	if sErr != nil {
		return sErr
	}

	zapLog = zap.New(logger.Core()).Sugar()
	return nil
}

func Info(message string, fields ...zap.Field) {
	zapLog.Info(message, fields)
}

func Fatal(message string, fields ...zap.Field) {
	zapLog.Fatal(message, fields)
}

func Error(message string, fields ...zap.Field) {
	zapLog.Error(message, fields)
}
