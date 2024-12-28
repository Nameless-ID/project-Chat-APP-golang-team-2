package controller

import (
	"net/http"
	"user-service/service"
	"user-service/util"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    service service.UserService
}

func NewUserController(service service.UserService) *UserController {
    return &UserController{service}
}

func (h *UserController) GetUserInfo(c *gin.Context) {
    userID := c.Param("user_id")
    user, err := h.service.GetUserInfo(userID)
    if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err.Error()))
			return
		}
        c.JSON(http.StatusInternalServerError, util.ErrorResponse("Error fetching user: "+err.Error()))
        return
    }
    c.JSON(http.StatusOK, user)
}

func (h *UserController) GetOnlineUsers(c *gin.Context) {
    users, err := h.service.GetOnlineUsers()
    if err != nil {
		if err.Error() == "users not found" {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err.Error()))
			return
		}
        c.JSON(http.StatusInternalServerError, util.ErrorResponse("Error fetching online user: "+err.Error()))
        return
    }
    c.JSON(http.StatusOK, users)
}
