package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang-program-structure/common/response"
)

// Status godoc
// @Summary Status
// @Description service status
// @Tags system
// @Produce json
// @Success 200 {object} response.StatusResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /v2/products/status [get]
func (config *Config) Status(c *gin.Context) {
	c.JSON(http.StatusOK, response.StatusResponse{Status: "ok"})
}
