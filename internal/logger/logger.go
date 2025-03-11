package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	config := NewConfig()
	l, _ := config.New()

	return l
}

func (c *Config) New() (*zap.Logger, error) {
	encoder := zapcore.NewConsoleEncoder(newEncoderConfig())

	return zap.New(zapcore.NewCore(
		encoder, zapcore.Lock(zapcore.AddSync(os.Stdout)), c.Level,
	), zap.Fields(zap.String(logIdKey, id()))), nil
}

func newEncoderConfig() zapcore.EncoderConfig {
	config := zap.NewProductionEncoderConfig()
	config.MessageKey = "message"
	config.TimeKey = "time"
	config.EncodeTime = zapcore.RFC3339TimeEncoder

	return config
}
