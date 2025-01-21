package main

import (
	userservice "github.com/ELi10-T/Scalable-E-Commerce-Platform/pkg/user-service"
)

// The core features of user service includes --
// handling user registration, authentication, and profile management.

func main() {
	server := userservice.RunUserService()
	server.Run(":8080")
}
