package controllers

import (
	"github.com/devdigvijay/gin-user-service/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func (u *UserController) Initalize(ge *gin.Engine) {
	var api *gin.RouterGroup = ge.Group("/user")

	api.POST("/save", u.userService.SaveUserInfomation())

}
