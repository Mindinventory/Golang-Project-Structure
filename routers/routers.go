package routers

import (
	apiControllerV1 "Golang-Project-Structure/controllers/api/v1"
	apiControllerV2 "Golang-Project-Structure/controllers/api/v2"
	"Golang-Project-Structure/middlewares"
	"github.com/gin-gonic/gin"
)

//SetupRouter function will perform all route operations
func SetupRouter() *gin.Engine {

	r := gin.Default()

	//Giving access to storage folder
	r.Static("/storage", "storage")

	//Giving access to template folder
	r.Static("/templates", "templates")
	r.LoadHTMLGlob("templates/*")

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	//API route for version 1
	v1 := r.Group("/api/v1")

	//If you want to pass your route through specific middlewares
	v1.Use(middlewares.UserMiddlewares())
	{
		v1.POST("user-list", apiControllerV1.UserList)
	}

	//API route for version 2
	v2 := r.Group("/api/v2")

	v2.POST("user-list", apiControllerV2.UserList)

	return r

}
