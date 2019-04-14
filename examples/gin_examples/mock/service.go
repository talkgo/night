package mock

import "ginexamples"

type UserAuthenticationProvider struct {
	LoginFnInvoked bool
	LoginFn        func(email string, password string) (*ginexamples.User, error)

	LogoutFnInvoked bool
	LogoutFn        func(sessionID string) error

	CheckAuthenticationFnInvoked bool
	CheckAuthenticationFn        func(sessionID string) (*ginexamples.User, error)
}

func (uAM *UserAuthenticationProvider) Login(email string, password string) (*ginexamples.User, error) {
	uAM.LoginFnInvoked = true
	return uAM.LoginFn(email, password)
}

func (uAM *UserAuthenticationProvider) Logout(sessionID string) error {
	uAM.LogoutFnInvoked = true
	return uAM.LogoutFn(sessionID)
}
func (uAM *UserAuthenticationProvider) CheckAuthentication(sessionID string) (*ginexamples.User, error) {
	uAM.CheckAuthenticationFnInvoked = true
	return uAM.CheckAuthenticationFn(sessionID)
}

type Authenticator interface {
	Hash(password string) (string, error)
	CompareHash(hashedPassword string, plainPassword string) (bool, error)
	SessionID() string
}

type AuthenticatorMock struct {
	HashFn        func(password string) (string, error)
	HashFnInvoked bool

	CompareHashFn        func(hashedPassword string, plainPassword string) error
	CompareHashFnInvoked bool

	SessionIDFn        func() string
	SessionIDFnInvoked bool
}

func (uAM *AuthenticatorMock) Hash(password string) (string, error) {
	uAM.HashFnInvoked = true
	return uAM.HashFn(password)
}

func (uAM *AuthenticatorMock) CompareHash(hashedPassword string, plainPassword string) error {
	uAM.CompareHashFnInvoked = true
	return uAM.CompareHashFn(hashedPassword, plainPassword)
}
func (uAM *AuthenticatorMock) SessionID() string {
	uAM.SessionIDFnInvoked = true
	return uAM.SessionIDFn()
}
