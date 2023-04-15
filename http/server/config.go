package server

import (
	"time"

	"go.uber.org/zap"
)

type Config struct {
	Addr           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	Logger         *zap.Logger
}

type Option func(*Config)

func WithLogger(l *zap.Logger) Option {
	return func(c *Config) {
		c.Logger = l
	}
}
