package middleware

import (
	"fmt"
	"time"

	"github.com/blendle/zapdriver"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"golang-program-structure/common/logging"
)

type RequestLogger struct {
	logger *zap.SugaredLogger
}

// The request logger takes the preconfigured logger used by the app logging
func NewRequestLogger(logger *zap.SugaredLogger) *RequestLogger {
	if logger == nil {
		logger = zap.NewNop().Sugar()
	}

	return &RequestLogger{
		logger,
	}
}

func (requestLogger *RequestLogger) Middleware(c *gin.Context) {
	start := time.Now()
	c.Next()

	sdReq := zapdriver.NewHTTP(c.Request, nil)

	sdReq.Status = c.Writer.Status()
	sdReq.Latency = fmt.Sprintf("%fs", time.Since(start).Seconds())

	l := requestLogger.logger
	if ctxRqId, ok := c.Value(logging.RequestIDKey).(string); ok {
		l = requestLogger.logger.With(zap.String(logging.LogRequestIDKey, ctxRqId))
	}
	l.Infow("Request ", zapdriver.HTTP(sdReq))
}
