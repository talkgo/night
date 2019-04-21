package postgres

import (
	"errors"
	"fmt"
	"ginexamples"

	"github.com/jinzhu/gorm"
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

func (u *UserRepository) Find(id string) (*ginexamples.User, error) {
	var user ginexamples.User

	db := u.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) FindByEmail(email string) (*ginexamples.User, error) {
	if email == "" {
		return &ginexamples.User{}, errors.New("not found")
	}
	return u.findBy("email", email)
}

func (u *UserRepository) FindBySessionID(sessionID string) (*ginexamples.User, error) {
	if sessionID == "" {
		return nil, errors.New("not found")
	}
	return u.findBy("session_id", sessionID)
}

func (u *UserRepository) findBy(key string, value string) (*ginexamples.User, error) {
	user := ginexamples.User{}

	db := u.db.Where(fmt.Sprintf("%s = ?", key), value)
	err := first(db, &user)

	return &user, err
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("resource not found")
	}
	return err
}

func (u *UserRepository) Update(user *ginexamples.User) error {
	return u.db.Save(user).Error
}
