package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang-program-structure/common/response"
)

func (config *Config) PanicRecovery(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		config.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Cause: "Something went wrong"})
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
