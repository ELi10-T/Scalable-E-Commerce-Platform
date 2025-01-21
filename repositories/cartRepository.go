package repositories

import (
	"github.com/ELi10-T/Scalable-E-Commerce-Platform/util"
	"github.com/ELi10-T/Scalable-E-Commerce-Platform/util/models"
	"gorm.io/gorm"
)

const (
	CARTTABLE = "carts_table"
	CARTITEMS = "cart_items"
	QUANTITY  = "quantity"
)

func NewCartRepository() *CartRepository {
	dbConn := util.InitDatabaseConn()
	return &CartRepository{
		dbConn: dbConn,
	}
}

type CartRepository struct {
	dbConn *gorm.DB
}

func (c *CartRepository) GetCart(id int) (*models.Cart, error) {
	// Use find for cart_items
	cart := &models.Cart{}
	tx := c.dbConn.Table(CARTTABLE).Take(cart, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	cartItems := []*models.CartItem{}
	tx = c.dbConn.Table(CARTITEMS).Find(&cartItems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	cart.Items = append(cart.Items, cartItems...)
	return cart, nil
}

func (c *CartRepository) AddToCart(id int, prodId int, quantity int) (*models.Cart, error) {
	cart, err := c.GetCart(id)
	if err != nil {
		return nil, err
	}

	for _, item := range cart.Items {
		if item.ProductId == prodId {
			item.Quantity += quantity
			tx := c.dbConn.Model(item).Update(QUANTITY, item.Quantity)
			if tx.Error != nil {
				return nil, tx.Error
			}
			return cart, nil
		}
	}

	item := &models.CartItem{
		CartId:    id,
		ProductId: prodId,
		Quantity:  quantity,
	}

	tx := c.dbConn.Table(CARTITEMS).Create(item)
	if tx.Error != nil {
		return nil, tx.Error
	}

	cart.Items = append(cart.Items, item)

	return cart, nil
}

func (c *CartRepository) CreateCart(userId int) error {
	cart := &models.Cart{
		UserId: userId,
	}
	tx := c.dbConn.Table(CARTTABLE).Create(cart)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
