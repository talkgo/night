package mock

import "ginexamples"

type UserRepository struct {
	StoreFn        func(user *ginexamples.User) error
	StoreFnInvoked bool

	UpdateFn        func(user *ginexamples.User) error
	UpdateFnInvoked bool

	FindFn        func(id string) (*ginexamples.User, error)
	FindFnInvoked bool

	FindByEmailFn        func(email string) (*ginexamples.User, error)
	FindByEmailFnInvoked bool

	FindBySessionIDFn        func(sessionID string) (*ginexamples.User, error)
	FindBySessionIDFnInvoked bool
}

func (uRM *UserRepository) Store(user *ginexamples.User) error {
	uRM.StoreFnInvoked = true
	return uRM.StoreFn(user)
}

func (uRM *UserRepository) Update(user *ginexamples.User) error {
	uRM.UpdateFnInvoked = true
	return uRM.UpdateFn(user)
}

func (uRM *UserRepository) Find(id string) (*ginexamples.User, error) {
	uRM.FindFnInvoked = true
	return uRM.FindFn(id)
}

func (uRM *UserRepository) FindByEmail(email string) (*ginexamples.User, error) {
	uRM.FindByEmailFnInvoked = true
	return uRM.FindByEmailFn(email)
}

func (uRM *UserRepository) FindBySessionID(sessionID string) (*ginexamples.User, error) {
	uRM.FindBySessionIDFnInvoked = true
	return uRM.FindBySessionIDFn(sessionID)
}
