package controller

import (
	"backend-app/service"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service.IUserService
}

type IUserController interface {
	GetUsersList(c *gin.Context)
}

func NewUserController(service service.IUserService) IUserController {
	return &userController{service}
}

func (controller *userController) GetUsersList(c *gin.Context) {
	user, err := controller.GetUsersService()
	if err != nil {
		slog.Error("Get user service error: ", err)
		os.Exit(1)
	}
	c.JSONP(http.StatusOK, gin.H{
		"user": user,
	})
}
