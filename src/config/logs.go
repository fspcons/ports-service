package config

import (
	"fmt"
	"go.uber.org/zap"
)

// MustNewLogger ensure logger must be provided
func MustNewLogger(serviceName string) *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("failed to start zap logger: %s\n", err.Error())
		panic(err.Error())
	}
	return logger.With(zap.String("service", serviceName))
}
