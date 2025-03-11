package logger

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logIdKey            string = "log_id"
	PostgresInstanceKey string = "postgres_instance"
	RedisInstanceKey    string = "redis_instance"
	LocalKey            string = "local"
)

func id() string {
	return uuid.NewString()
}

func Postgres(name string) zapcore.Field {
	return zap.String(PostgresInstanceKey, name)
}

func Redis(name string) zapcore.Field {
	return zap.String(RedisInstanceKey, name)
}

func Local(name string) zapcore.Field {
	return zap.String(LocalKey, name)
}
