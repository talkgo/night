package userservice

import (
	"ginexamples"

	"github.com/pkg/errors"
)

type UserService struct {
	r ginexamples.UserRepository
}

// New returns the UserService.
func New(userRepository ginexamples.UserRepository) *UserService {
	return &UserService{
		r: userRepository,
	}
}

func (uS *UserService) CreateUser(user *ginexamples.User, password string) (*ginexamples.User, error) {
	_, err := uS.r.FindByEmail(user.Email)
	if err == nil {
		return &ginexamples.User{}, err
	}

	if len(password) < 8 {
		return &ginexamples.User{}, errors.New("password too short")
	}

	// if err != nil {
	// 	return &ginexamples.User{}, errors.Wrap(err, "error hashing password")
	// }

	user.PasswordHash = password

	err = uS.r.Store(user)
	if err != nil {
		return &ginexamples.User{}, errors.Wrap(err, "error storing user")
	}
	return user, nil
}
