package cartservice

import (
	"github.com/ELi10-T/Scalable-E-Commerce-Platform/repositories"
	"github.com/gin-gonic/gin"
)

type CartService struct {
	cartRepo *repositories.CartRepository
}

func RunCartService() *gin.Engine {
	server := gin.Default()
	cartService := CartService{}
	cartService.cartRepo = repositories.NewCartRepository()

	cartServer := server.Group("/carts")

	cartServer.GET("/:id", cartService.getCart)
	cartServer.POST("/:id", cartService.addToCart)

	return server
}
