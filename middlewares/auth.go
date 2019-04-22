package middlewares

import "github.com/gin-gonic/gin"

//Middleware functions
func UserMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Code for middlewares

		c.Next()
	}
}
