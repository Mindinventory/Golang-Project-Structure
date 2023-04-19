package helpers

import (
	"github.com/gin-gonic/gin"

	"golang-program-structure/common/response"
)

func RespondWithError(c *gin.Context, code int, cause string) {
	c.JSON(code, response.ErrorResponse{
		Code:  code,
		Cause: cause,
	})
}
