package userservice

import (
	"github.com/ELi10-T/Scalable-E-Commerce-Platform/util/models"
	"github.com/gin-gonic/gin"
)

const (
	TABLENAME = "user_table"
)

func (u *UserService) CreateUser(ctx *gin.Context) {
	createUser := models.User{}

	if err := ctx.BindJSON(&createUser); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := u.userRepo.CreateUser(&createUser)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(201, createUser)
}

func (u *UserService) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	getUser, err := u.userRepo.GetUser(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, getUser)
}

func (u *UserService) Login(ctx *gin.Context) {

	// TODO: 
	// 1. Hash the password
	// 2. Send back JWT -- learn about this

	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	getUser, err := u.userRepo.GetUserBasedOnQuery("username=?", request.Username)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if getUser.Password != request.Password {
		ctx.JSON(401, gin.H{
			"response": "unauthorized",
		})
		return
	}

	ctx.JSON(200, getUser)

}
