package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

const bearerPrefix = "Bearer "

type UserTokenData struct {
	Id       uuid.UUID `json:"id" bson:"id"`
	Email    string    `json:"email" bson:"email"`
	UserType string    `json:"user_type" bson:"user_type"`
	Scope    string    `json:"scope"`
}

type AdminTokenData struct {
	Id   uuid.UUID `json:"id" bson:"id"`
	Type string    `json:"type" bson:"type"`
}

type AuthMiddlewareConfig struct {
	RequestAuth
}

type RequestAuth struct {
	logger *zap.SugaredLogger
}

func abortAuthenticationFailed(c *gin.Context) {
	// Send back the Unauthorized message
	c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "cause": "Unauthorized"})
	// Don't go any further down the request chain
	c.Abort()
}

// The request logger takes the preconfigured logger used by the app logging
func NewRequestAuth(logger *zap.SugaredLogger) *RequestAuth {
	return &RequestAuth{
		logger,
	}
}

func (ra *RequestAuth) Middleware(c *gin.Context) {

	// check if the user is authenticated
	authPassed := ra.BearerAuth(c)

	// not authenticated, no access
	if !authPassed {
		abortAuthenticationFailed(c)
		return
	}
	c.Next()
}

func (ra *RequestAuth) BearerAuth(c *gin.Context) bool {

	authHeaderValue := c.Request.Header.Get("Authorization")

	if !strings.HasPrefix(authHeaderValue, bearerPrefix) {
		ra.logger.Errorf("No bearer token found in Authorization header")
		return false
	}

	token := strings.TrimPrefix(authHeaderValue, bearerPrefix)
	if len(token) == 0 {
		ra.logger.Error("No bearer token found in Authorization header")
		return false
	}

	return true
}
