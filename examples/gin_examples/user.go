package ginexamples

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	SessionID    string `gorm:"not null"`
}

type UserRepository interface {
	Store(user *User) error
	Update(user *User) error
	Find(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindBySessionID(sessionID string) (*User, error)
}

type UserService interface {
	CreateUser(u *User, password string) (*User, error)
	GetUser(id string) (*User, error)
	UserAuthenticationProvider
}

type UserAuthenticationProvider interface {
	Login(email string, password string) (*User, error)
	Logout(sessionID string) error
	CheckAuthentication(sessionID string) (*User, error)
}
