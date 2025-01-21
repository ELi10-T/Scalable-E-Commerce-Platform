package cartservice

import (
	"net/http"
	"strconv"

	"github.com/ELi10-T/Scalable-E-Commerce-Platform/util/models"
	"github.com/gin-gonic/gin"
)



func (c *CartService) getCart(ctx *gin.Context) {
	// TODO: 1. Get all cart items first
	//  2. Take all into the cart
	//  3. 

	id := ctx.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 0)
	cart, err := c.cartRepo.GetCart(int(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": cart,
	})
}

func (c *CartService) addToCart(ctx *gin.Context) {

	// add the logic of ignoring, if product is greater

	cartItem := &models.CartItem{}
	if err := ctx.BindJSON(cartItem); err != nil {
		ctx.JSON(500, gin.H{
			"error":err,
		})
	}

	cart, err := c.cartRepo.AddToCart(cartItem.CartId, cartItem.ProductId, cartItem.Quantity)

	if err != nil {
		ctx.JSON(500, gin.H{
			"Error": err,
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": cart,
	})
}

func (c *CartService) addCart(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 0)
	err := c.cartRepo.CreateCart(int(idInt))
	if err != nil {
		ctx.JSON(500, err)
	}

	ctx.JSON(200, nil)
}