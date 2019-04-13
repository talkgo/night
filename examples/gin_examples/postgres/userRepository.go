package postgres

import (
	"errors"
	"fmt"
	"ginexamples"

	"github.com/jinzhu/gorm"
)

const (
	// ErrNotFound is returned when a resource cannot be found
	// in the database.
	ErrNotFound modelError = "UserRepository: resource not found"
)

type UserRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Stroe creates a user record in the table
func (u *UserRepository) Store(user *ginexamples.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepository) FindByEmail(email string) (*ginexamples.User, error) {
	if email == "" {
		return &ginexamples.User{}, errors.New("not found")
	}
	return u.findBy("email", email)
}

func (u *UserRepository) findBy(key string, value string) (*ginexamples.User, error) {
	user := ginexamples.User{}

	db := u.db.Where(fmt.Sprintf("%s = ?", key), value)
	err := first(db, &user)

	return &user, err
}

// first will query using the provided gorm.DB and it will
// get the first item returned and place it into dst. If
// nothing is found in the query, it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

func (u *UserRepository) Update(user *ginexamples.User) error {
	return u.db.Save(user).Error
}

type modelError string

func (e modelError) Error() string {
	return string(e)
}
