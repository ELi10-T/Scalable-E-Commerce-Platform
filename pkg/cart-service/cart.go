package cartservice

import "github.com/gin-gonic/gin"

func (c *CartService) getCart(ctx *gin.Context) {
	id := ctx.Param("id")
	c.cartRepo.GetCart()



}

func (c *CartService) addToCart(ctx *gin.Context) {
	// id := ctx.Param("id")

	// check whether the cart exists
	// if exists, check the product to be added
		// -- if new, then add product
		// -- if old, increase quantity
}