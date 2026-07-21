package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	logOnce sync.Once
	logger  *zap.Logger

	retrieveEnvFunc = func() string {
		return os.Getenv(EnvDeploymentKey)
	}
)

func SetLogEntry(zlg *zap.Logger) {
	if zlg == nil {
		fmt.Println("entry logger set is unavailable, skipping logger setup")
		return
	}

	// cache the logger
	buffered := logger
	// Set the logger by the given zap.Logger
	logger = zlg
	// Flush any buffered entry entries
	if buffered != nil {
		if err := buffered.Sync(); err != nil {
			logger.Debug("Flush the buffered entry entries", zap.Error(err))
		}
	}
}

// DeploymentEnvWithFunc sets the function to retrieve the environment value for logger initialization.
//
// You can provide a custom function to retrieve the environment value, for example, from a configuration file or an external service.
//
// The function should return a string representing the environment (e.g., "development", "production").
//
// Usage:
//
//	fn := func() string {
//	 return os.Getenv("DEPLOYMENT_ENV")
//	}
//	// Sets the function to retrieve with #fn
//	DeploymentEnvWithFunc(fn)
//
// If the function is nil or returns an empty string, the environment setup will be skipped.
func DeploymentEnvWithFunc(fn func() string) {
	if fn == nil || fn() == "" {
		fmt.Println("environment function is unavailable, skipping environment setup")
		return
	}
	retrieveEnvFunc = fn
}

func NewEntry() *zap.Logger {
	logOnce.Do(func() {
		if logger != nil {
			// If logger is already initialized, skip re-initialization
			return
		}
		// Renew the logger default
		val := retrieveEnvFunc()
		switch val {
		case "development", "dev":
			// Initialize logger
			config := zap.NewDevelopmentConfig()
			config.Encoding = "json"
			dev, err := config.Build()
			if err != nil {
				logger = zap.NewExample()
				logger.Debug("failed to initialize production logger", zap.Error(err))
			} else {

				logger = dev
			}
		default:
			// Consider production environment
			// Initialize logger
			prod, err := zap.NewProduction()
			if err != nil {
				logger = zap.NewExample()
				logger.Debug("failed to initialize production logger", zap.Error(err))
			} else {
				logger = prod
			}
		}
	})

	return logger.With(
		zap.Time(KeyTimestamp, time.Now()),
		zap.String(KeyEnvironment, retrieveEnvFunc()))
}
