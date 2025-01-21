package models

type User struct {
	Id		 int	 
	Name     string    `json:"name" binding:"required" gorm:"column:username"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required" gorm:"column:password_hash"`
}