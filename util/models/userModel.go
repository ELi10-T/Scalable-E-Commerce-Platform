package models

// From Frontend, what is expected to be sent?
// Name, Email, UserName, Password

// What are all the things that the user is expected to do with his profile?
// 1. Create a profile
// 2. Login
// 3. Logout
// 4. Update profile
// 5. Delete profile

type User struct {
	Id       int
	Name     string `json:"name" binding:"required" gorm:"column:username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required" gorm:"column:password_hash"`
}
