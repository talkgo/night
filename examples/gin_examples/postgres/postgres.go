package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (c DBConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s "+
			"sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s sslmode=disable", c.Host, c.Port, c.User,
		c.Password, c.Name)
}

type Repository struct {
	UserRepository *UserRepository
}

func Initialize(c DBConfig) *Repository {
	db, err := gorm.Open("postgres", c.ConnectionInfo())
	if err != nil {
		panic(err)
	}

	return &Repository{
		UserRepository: newUserRepository(db),
	}
}
