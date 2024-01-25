package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

type LogContextKey string

const loggerCtxKey LogContextKey = "logger_ctx_key"

type loggerContextData struct {
	fields Fields
}

func WithFields(ctx context.Context, fields Fields) context.Context {
	if len(fields) == 0 {
		return ctx
	}

	data := extractLoggerData(ctx)
	for key, value := range fields {
		data.fields[key] = value
	}
	return context.WithValue(ctx, loggerCtxKey, data)
}
func extractLoggerData(ctx context.Context) loggerContextData {
	logData := loggerContextData{
		fields: make(map[string]interface{}),
	}

	existingLogData, ok := ctx.Value(loggerCtxKey).(loggerContextData)
	if ok {
		for key, val := range existingLogData.fields {
			logData.fields[key] = val
		}
	}
	return logData
}

func FromContext(ctx context.Context) *logrus.Entry {
	return logrus.WithFields(logrus.Fields(extractLoggerData(ctx).fields))
}
