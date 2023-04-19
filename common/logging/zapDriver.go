package logging

import (
	"github.com/blendle/zapdriver"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	RequestIDKey    = "X-Request-ID"
	LogRequestIDKey = "req-id"
)

// See: https://github.com/evry/go-negroni-zap-middleware

func GetLogger() *zap.SugaredLogger {
	// The sugared logger allows untyped arguments, similar to "fmt.Printf"
	return zap.S()
}

func GetRequestLogger(c *gin.Context) *zap.SugaredLogger {
	// The sugared logger allows untyped arguments, similar to "fmt.Printf"
	l := zap.S()
	if c != nil {
		if ctxRqId, ok := c.Value(RequestIDKey).(string); ok {
			l = l.With(zap.String(LogRequestIDKey, ctxRqId))
		}
	}
	return l
}

func SetupLogger(environment string) {
	// By default, use the kubernetes specific production StackDriver logger
	logger, err := zapdriver.NewProduction()

	// In local mode, use the console logger
	if environment == "local" {
		logger, err = zap.NewDevelopment()
	}
	// In test mode, don't output warnings
	if environment == "test" {
		testConfig := zap.NewDevelopmentConfig()
		testConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		logger, err = testConfig.Build()
	}
	if err != nil {
		panic(err)
	}

	// Replace the global Zap logger with the logger we just configured
	zap.ReplaceGlobals(logger)

	GetLogger().Info("Logger configured")
}
