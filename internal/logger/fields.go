package logger

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	_LogIdKey            string = "log_id"
	_PostgresInstanceKey string = "postgres_instance"
	_RedisInstanceKey    string = "redis_instance"
	_LocalKey            string = "local"
	_ErrDetailsKey       string = "error_details"
	_ServerDetailsKey    string = "server_details"
)

func id() string {
	return uuid.NewString()
}

func Postgres(name string) zapcore.Field {
	return zap.String(_PostgresInstanceKey, name)
}

func Redis(name string) zapcore.Field {
	return zap.String(_RedisInstanceKey, name)
}

func Local(name string) zapcore.Field {
	return zap.String(_LocalKey, name)
}

func ErrDetails(err error) zapcore.Field {
	return zap.NamedError(_ErrDetailsKey, err)
}

func Server(addr string) zapcore.Field {
	return zap.String(_ServerDetailsKey, addr)
}
