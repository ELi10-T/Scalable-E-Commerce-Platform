package models

type Cart struct {
	Id     int `json:"id"`
	UserId int `json:"user_id" gorm:"user_id"`
	Items []*CartItem `json:"items" gorm:"-"`
}

type CartItem struct {
	Id int `json:"id"`
	CartId int `json:"cart_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
