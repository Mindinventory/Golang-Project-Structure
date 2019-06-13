package middlewares

import "github.com/gin-gonic/gin"

/*
UserMiddlewares function to add auth
*/
func UserMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Code for middlewares

		c.Next()
	}
}
