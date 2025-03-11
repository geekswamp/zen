package logger

import "go.uber.org/zap/zapcore"

type Config struct {
	Level zapcore.LevelEnabler
}

func NewConfig() Config {
	return Config{Level: zapcore.Level(0)}
}
