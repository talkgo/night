package userservice

import (
	"ginexamples"
	"ginexamples/pkg/auth"

	"github.com/pkg/errors"
)

type UserService struct {
	r ginexamples.UserRepository
	a Authenticator
}

type Authenticator interface {
	Hash(password string) (string, error)
	CompareHash(hashedPassword string, plainPassword string) error
	SessionID() string
}

// New returns the UserService.
func New(userRepository ginexamples.UserRepository) *UserService {
	return &UserService{
		r: userRepository,
		a: &auth.Authenticator{},
	}
}

func (uS *UserService) CreateUser(user *ginexamples.User, password string) (*ginexamples.User, error) {
	_, err := uS.r.FindByEmail(user.Email)
	if err == nil {
		return &ginexamples.User{}, errors.New("email already exists")
	}

	if len(password) < 8 {
		return &ginexamples.User{}, errors.New("password too short")
	}

	hashedPassword, err := uS.a.Hash(password)
	if err != nil {
		return &ginexamples.User{}, errors.Wrap(err, "error hashing password")
	}

	user.PasswordHash = hashedPassword
	user.SessionID = uS.a.SessionID()

	err = uS.r.Store(user)
	if err != nil {
		return &ginexamples.User{}, errors.Wrap(err, "error storing user")
	}
	return user, nil
}

func (uS *UserService) Login(email string, password string) (*ginexamples.User, error) {
	user, err := uS.r.FindByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "error finding user by email")
	}

	err = uS.a.CompareHash(user.PasswordHash, password)
	if err != nil {
		return nil, errors.Wrap(err, "error comparing hash")
	}

	user.SessionID = uS.a.SessionID()
	err = uS.r.Update(user)
	if err != nil {
		return nil, errors.Wrap(err, "error updating sessionID")
	}

	return user, nil
}

func (uS *UserService) Logout(sessionID string) error {
	user, err := uS.r.FindBySessionID(sessionID)
	if err != nil {
		return errors.Wrap(err, "error finding by sessionID")
	}

	user.SessionID = ""
	uS.r.Update(user)

	return nil
}

func (uS *UserService) CheckAuthentication(sessionID string) (*ginexamples.User, error) {
	user, err := uS.r.FindBySessionID(sessionID)
	if err != nil {
		return nil, errors.Wrap(err, "error finding by sessionID")
	}

	return user, nil
}

func (uS *UserService) GetUser(id string) (*ginexamples.User, error) {
	return uS.r.Find(id)
}
