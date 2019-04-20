package userservice

import (
	"ginexamples"
	"ginexamples/pkg/mock"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var r = mock.UserRepository{
	StoreFnInvoked: false,
	StoreFn: func(user *ginexamples.User) error {
		if user.Email == "fail@mail.com" {
			return errors.New("error storing user")
		}
		return nil
	},
	UpdateFnInvoked: false,
	UpdateFn: func(user *ginexamples.User) error {
		return nil
	},
	FindFnInvoked: false,
	FindFn: func(id string) (*ginexamples.User, error) {
		if id == "1" {
			return &ginexamples.User{Model: gorm.Model{ID: 1}}, nil
		}
		return nil, errors.New("not found")
	},
	FindByEmailFnInvoked: false,
	FindByEmailFn: func(email string) (*ginexamples.User, error) {
		if email == "existing@mail.com" {
			return &ginexamples.User{Model: gorm.Model{ID: 1}, PasswordHash: "password"}, nil
		}
		return nil, errors.New("not found")
	},
	FindBySessionIDFnInvoked: false,
	FindBySessionIDFn: func(sessionID string) (*ginexamples.User, error) {
		if sessionID == "sessionID-0-0-0" {
			return &ginexamples.User{Model: gorm.Model{ID: 1}}, nil
		}
		return nil, errors.New("not found")
	},
}
var a = mock.AuthenticatorMock{
	HashFnInvoked: false,
	HashFn: func(password string) (string, error) {
		if password == "longpasswordbuthashfailed" {
			return "", errors.New("Error hashing password")
		}
		return password, nil
	},
	CompareHashFnInvoked: false,
	CompareHashFn: func(hashedPassword string, plainPassword string) error {
		if hashedPassword != plainPassword {
			return errors.New("incorrect password")
		}
		return nil
	},
	SessionIDFnInvoked: false,
	SessionIDFn: func() string {
		return "sessionID-0-0-0"
	},
}

var us = &UserService{&r, &a}

func TestNew(t *testing.T) {
	t.Parallel()
	u := New(&r)
	assert.Equal(t, &r, u.r, "repository does not match")
	assert.NotNil(t, u.a, "no authenticator")
}

func TestUserService_CreateUser(t *testing.T) {
	var testCases = []struct {
		name     string
		email    string
		password string
		error    bool
	}{
		{"user exists", "existing@mail.com", "password", true},
		{"password too short", "new@mail.com", "pass", true},
		{"error hashing password", "new@mail.com", "longpasswordbuthashfailed", true},
		{"error storing user", "fail@mail.com", "longpassword", true},
		{"new user", "new@mail.com", "longpassword", false},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			defer func() { r.FindByEmailFnInvoked = false }()
			defer func() { a.SessionIDFnInvoked = false }()
			defer func() { r.StoreFnInvoked = false }()
			defer func() { a.HashFnInvoked = false }()

			user, err := us.CreateUser(&ginexamples.User{Email: v.email}, v.password)
			if v.error {
				assert.NotNil(t, err, "did not fail to create existing user")
				if err.Error() != "error storing user: error storing user" {
					assert.False(t, r.StoreFnInvoked, "Store was invoked")
				}
				assert.Empty(t, user, "did not return empty user")
				return
			}

			assert.Nil(t, err, "should not fail to create user")
			assert.True(t, r.FindByEmailFnInvoked, "FindByEmail was not invoked")
			assert.True(t, a.HashFnInvoked, "Hash was not invoked")
			assert.Equal(t, v.password, user.PasswordHash, "did not hash password")
			assert.True(t, a.SessionIDFnInvoked, "SessionID was not invoked")
			assert.Equal(t, "sessionID-0-0-0", user.SessionID, "sessionID was not set")
			assert.True(t, r.StoreFnInvoked, "Store was not invoked")
		})
	}
}
func TestUserService_Login(t *testing.T) {
	var testCases = []struct {
		name     string
		email    string
		password string
		error    bool
	}{
		{"new user", "new@mail.com", "password", true},
		{"incorrect password", "existing@mail.com", "pass2", true},
		{"user exists", "existing@mail.com", "password", false},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			defer func() { r.FindByEmailFnInvoked = false }()
			defer func() { r.UpdateFnInvoked = false }()
			defer func() { a.CompareHashFnInvoked = false }()
			defer func() { a.SessionIDFnInvoked = false }()

			user, err := us.Login(v.email, v.password)
			if v.error {
				assert.NotNil(t, err, "did not fail to log in user")
				assert.Empty(t, user, "did not return empty user")
				return
			}
			assert.Nil(t, err, "should not fail to login userService")
			assert.True(t, r.FindByEmailFnInvoked, "FindByEmail not invoked")
			assert.True(t, r.UpdateFnInvoked, "Update not invoked")
			assert.True(t, a.CompareHashFnInvoked, "CompareHash not invoked")
			assert.True(t, a.SessionIDFnInvoked, "SessionID not invoked")
			assert.Equal(t, "sessionID-0-0-0", user.SessionID, "did not update sessionID")
		})
	}
}

func TestUserService_Logout(t *testing.T) {
	var testCases = []struct {
		name      string
		sessionID string
		error     bool
	}{
		{"unknown session", "someOtherSession", true},
		{"user exists", "sessionID-0-0-0", false},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			defer func() { r.FindBySessionIDFnInvoked = false }()
			defer func() { r.UpdateFnInvoked = false }()

			err := us.Logout(v.sessionID)
			if v.error {
				assert.NotNil(t, err, "did not fail")
				assert.True(t, r.FindBySessionIDFnInvoked, "FindBySessionID was not invoked")
				assert.False(t, r.UpdateFnInvoked, "UpdateSessionIDFn was invoked")
				return
			}
			assert.Nil(t, err, "caused error logging out user")
			assert.True(t, r.FindBySessionIDFnInvoked, "FindBySessionID was not invoked")
			assert.True(t, r.UpdateFnInvoked, "UpdateSessionID was not invoked")
		})
	}
}
func TestUserService_CheckAuthentication(t *testing.T) {
	var testCases = []struct {
		name      string
		sessionID string
		error     bool
	}{
		{"unknown session", "someOtherSession", true},
		{"user exists", "sessionID-0-0-0", false},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			defer func() { r.FindBySessionIDFnInvoked = false }()

			user, err := us.CheckAuthentication(v.sessionID)
			if v.error {
				assert.NotNil(t, err, "did not fail")
				assert.Empty(t, user, "should not contain user")
				assert.True(t, r.FindBySessionIDFnInvoked, "UpdateSessionIDFn was invoked")
				return
			}
			assert.Nil(t, err, "caused error logging out user")
			assert.NotEmpty(t, user, "should contain user")
			assert.True(t, r.FindBySessionIDFnInvoked, "FindBySessionID was not invoked")
		})
	}
}
func TestUserService_GetUser(t *testing.T) {
	var testCases = []struct {
		name  string
		id    string
		error bool
	}{
		{"user exists", "1", false},
		{"new user", "2", true},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			defer func() { r.FindFnInvoked = false }()

			user, err := us.GetUser(v.id)
			if v.error {
				assert.NotNil(t, err, "did not fail to get non-existing user")
				assert.Empty(t, user, "did not return empty user")
				assert.True(t, r.FindFnInvoked, "Find was not invoked")
				return
			}
			assert.Nil(t, err, "failed to get existing user")
			assert.NotEmpty(t, user, "returned non-empty user")
			assert.True(t, r.FindFnInvoked, "Find was not invoked")
		})
	}
}
