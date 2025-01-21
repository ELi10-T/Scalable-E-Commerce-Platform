package models

type Cart struct {
	Id     int
	UserId string `json:"user_id" gorm:"user_id"`
	Items []CartItem `json:"items" gorm:"-"`
}

type CartItem struct {
	Id string 
	CartId string 
	ProductId string
	Quantity  int
}
