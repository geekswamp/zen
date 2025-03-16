package logger

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logIdKey            string = "log_id"
	postgresInstanceKey string = "postgres_instance"
	redisInstanceKey    string = "redis_instance"
	localKey            string = "local"
	errDetailsKey       string = "error_details"
)

func id() string {
	return uuid.NewString()
}

func Postgres(name string) zapcore.Field {
	return zap.String(postgresInstanceKey, name)
}

func Redis(name string) zapcore.Field {
	return zap.String(redisInstanceKey, name)
}

func Local(name string) zapcore.Field {
	return zap.String(localKey, name)
}

func ErrDetails(err error) zapcore.Field {
	return zap.NamedError(errDetailsKey, err)
}
