package userservice

import (
	"os"
	"strconv"
	"time"

	"github.com/ELi10-T/Scalable-E-Commerce-Platform/util/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	TABLENAME = "user_table"
)

// This API will be called when the user wants to create a new account
func (u *UserService) CreateUser(ctx *gin.Context) {
	createUser := models.User{}

	if err := ctx.BindJSON(&createUser); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// before creating the user, we need to check if the user already exists
	// assuming email would be unique
	existingUser, err := u.userRepo.GetUserBasedOnQuery("email=?", createUser.Email)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if existingUser != nil {
		ctx.JSON(400, gin.H{
			"error": "User already exists",
		})
		return
	}

	// before creating the user, we need to hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	createUser.Password = string(hashedPassword)

	err = u.userRepo.CreateUser(&createUser)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(201, createUser)
}

// This API will be called when the user wants to get his profile,
// But how does he send the id?
func (u *UserService) GetUser(ctx *gin.Context) {
	// TODO: Add JWT verification
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	if claims.Claims.(jwt.MapClaims)["exp"].(int64) < time.Now().Unix() {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	userId := claims.Claims.(jwt.MapClaims)["sub"].(int64)

	getUser, err := u.userRepo.GetUser(strconv.FormatInt(userId, 10))
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, getUser)
}

// This API will be called when the user wants to login to his account
func (u *UserService) Login(ctx *gin.Context) {

	// TODO:
	// 1. Hash the password
	// 2. Send back JWT -- learn about this

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	getUser, err := u.userRepo.GetUserBasedOnQuery("email=?", request.Email)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	// compare hash and password
	err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(request.Password))
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	// if successful, we need to generate a JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": getUser.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}

// where exactly is JWT used?
// JWT is used to authenticate the user and authorize the user to access the resources
// JWT is a token that is used to authenticate the user and authorize the user to access the resources
// JWT is a token that is used to authenticate the user and authorize the user to access the resources
