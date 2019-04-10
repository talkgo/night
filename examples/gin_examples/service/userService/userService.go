package userService

import "gin_examples"

type UserService struct {
	r gin_examples.UserRepository
}

func New(userRepository gin_examples.UserRepository) *UserService {
	return &UserService{
		r: userRepository,
	}
}
