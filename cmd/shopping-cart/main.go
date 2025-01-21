package main

import cartservice "github.com/ELi10-T/Scalable-E-Commerce-Platform/pkg/cart-service"

// shopping cart service
func main() {
	server := cartservice.RunCartService()
	server.Run(":8080")
}