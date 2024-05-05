package controller

import (
	"backend-app/service"
	"net/http"

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
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
	})
}
