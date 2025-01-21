package repositories

import (
	"github.com/ELi10-T/Scalable-E-Commerce-Platform/util"
	"github.com/ELi10-T/Scalable-E-Commerce-Platform/util/models"
	"gorm.io/gorm"
)

const (
	CARTTABLE = "cart_table"
	CARTITEMS = "cart_items"
)

func NewCartRepository() *CartRepository {
	cartTable := util.InitDatabaseConn().Table(CARTTABLE)
	cartItems := util.InitDatabaseConn().Table(CARTITEMS)
	return &CartRepository{
		cartTable: cartTable,
		cartItemsTable: cartItems,
	}
}

type CartRepository struct {
	cartTable *gorm.DB
	cartItemsTable *gorm.DB
}

func (c *CartRepository) GetCart(id string) (*models.Cart, error) {
	// Use find for cart_items
	
}

func (c *CartRepository) AddCart() {

}

func (c *CartRepository) AddToCart() {

}