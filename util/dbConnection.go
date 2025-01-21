package util

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Need to use GORM to connect to DB

const (
	Host     = "localhost"
	Port     = "5432"
	User     = "postgres"
	Password = "test"
)

func InitDatabaseConn() *gorm.DB {

	// need to understand this piece of code

	// This line is constructing the Data Source Name (DSN) which is a connection string for connecting to the PostgreSQL database. The connection string contains the necessary information for establishing a connection, such as the host, port, user, password, and other options like SSL mode.
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", Host, Port, User, Password)
	
	// Here, postgres.Open(dsn) tells GORM that the database you want to connect to is PostgreSQL and provides the connection string (dsn). This is how GORM knows it needs to use the PostgreSQL driver to interact with the database.
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("Unable to initiate the database")
	}

	return dbInstance
}
