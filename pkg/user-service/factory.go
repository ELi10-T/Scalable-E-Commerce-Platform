package userservice

import (
	"github.com/ELi10-T/Scalable-E-Commerce-Platform/repositories"
	"github.com/gin-gonic/gin"
)

// Enable db connectivity here
// Use Postgresql

type UserService struct {
	userRepo *repositories.UserRepository
	// logger  -- need to have
}

func RunUserService() *gin.Engine {
	userService := UserService{}
	userService.userRepo = repositories.NewUserRepository()
	server := gin.Default()

	// initiating the server
	userServer := server.Group("/users")
	userServer.POST("/register", userService.CreateUser)
	userServer.GET("/:id", userService.GetUser)
	userServer.POST("/login", userService.Login)

	return server
}
