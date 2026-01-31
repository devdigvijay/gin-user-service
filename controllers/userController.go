package controllers

import (
	"github.com/devdigvijay/gin-user-service/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func (u *UserController) Initialize(ge *gin.Engine) {
	var api *gin.RouterGroup = ge.Group("user")

	api.GET("/", u.userService.GetUserInfomation())
	api.POST("/save", u.userService.SaveUserInfomation())

}
