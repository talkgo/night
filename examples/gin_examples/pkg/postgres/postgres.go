package postgres

import (
	"fmt"
	"ginexamples"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConfig contains the environment varialbes requirements to initialize postgres.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (c DBConfig) connectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s "+
			"sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s sslmode=disable", c.Host, c.Port, c.User,
		c.Password, c.Name)
}

// Repository contains information for every repositories.
type Repository struct {
	UserRepository *UserRepository
	db             *gorm.DB
}

// Initialize the postgres database.
func Initialize(c DBConfig) *Repository {
	db, err := gorm.Open("postgres", c.connectionInfo())
	if err != nil {
		panic(err)
	}

	return &Repository{
		UserRepository: newUserRepository(db),
		db:             db,
	}
}

// AutoMigrate will attempt to automatically migrate all tables
func (r *Repository) AutoMigrate() error {
	err := r.db.AutoMigrate(&ginexamples.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
