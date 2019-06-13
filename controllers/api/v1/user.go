package v1

import (
	u "GoProject/apiHelpers"
	v1s "GoProject/services/api/v1"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//UserList function will give you the list of users
func UserList(c *gin.Context) {
	var userService v1s.UserService

	//decode the request body into struct and failed if any error occur
	err := json.NewDecoder(c.Request.Body).Decode(&userService.User)
	if err != nil {
		u.Respond(c.Writer, u.Message(1, "Invalid request"))
		return
	}

	//call service
	resp := userService.UserList()

	//return response using api helper
	u.Respond(c.Writer, resp)

}
