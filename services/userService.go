package services

import (
	"log"
	"net/http"

	"github.com/devdigvijay/gin-user-service/models/requests"
	"github.com/devdigvijay/gin-user-service/models/responses"
	"github.com/devdigvijay/gin-user-service/utils"
	"github.com/gin-gonic/gin"
)

type UserService struct{}

func (us *UserService) SaveUserInfomation() gin.HandlerFunc {

	return func(context *gin.Context) {
		var request requests.CreateUserRequest
		if error := context.ShouldBindJSON(&request); error != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": error.Error(),
			})
			return
		}

		log.Println(utils.ToJson(request))
		context.JSON(http.StatusOK, responses.CreateUserResponse{
			Id:       "101",
			UserName: request.UserName,
			IsActive: true,
		})
	}
}

func (us *UserService) GetUserInfomation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responses.CreateUserResponse{
			Id:       "101",
			UserName: "digv1705",
		})
	}
}
