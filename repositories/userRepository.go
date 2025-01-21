package repositories

import (
	"github.com/ELi10-T/Scalable-E-Commerce-Platform/util"
	"github.com/ELi10-T/Scalable-E-Commerce-Platform/util/models"
	"gorm.io/gorm"
)

const (
	USERTABLE = "user_table"
)

func NewUserRepository() *UserRepository {
	dbConn := util.InitDatabaseConn().Table(USERTABLE)
	return &UserRepository{
		dbConn: dbConn,
	}
}

type UserRepository struct {
	dbConn *gorm.DB
}

func (u *UserRepository) CreateUser(user *models.User) (error) {
	tx := u.dbConn.Create(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *UserRepository) GetUser(id string) (*models.User, error) {
	getUser := &models.User{}
	tx := u.dbConn.Take(getUser, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return getUser, nil
}

func (u *UserRepository) GetUserBasedOnQuery(query string, value string) (*models.User, error) {
	getUser := &models.User{}
	tx := u.dbConn.Take(getUser, query, value)
	
	if tx.Error != nil {
		return nil, tx.Error
	}

	return getUser, nil
}