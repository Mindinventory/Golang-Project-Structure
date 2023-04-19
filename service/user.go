package service

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/opencontainers/runc/libcontainer/utils"

	"golang-program-structure/common/helpers"
	"golang-program-structure/repository"
)

// GetUser godoc
// @Summary Get a user
// @Description get user by user id
// @Tags user
// @Param userId path string true "User ID"
// @Accept json
// @Produce json
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /user [get]
func (config *Config) GetUser(c *gin.Context) {

	userId, err := uuid.FromString(c.Param("userId"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := repository.GetUserById(context.Background(), &userId)
	if err != nil {
		config.Logger.Errorf("Error while searching: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "internal server error")
		return
	}
	if user.ID.IsNil() {
		helpers.RespondWithError(c, http.StatusNotFound, "user not found")
		return
	}

	err = utils.WriteJSON(c.Writer, c)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, user)
}
