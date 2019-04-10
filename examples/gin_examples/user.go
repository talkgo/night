package gin_examples

type User struct {
}

type UserRepository interface {
	// Find(id string) (error, User)
}

type UserService interface {
	// CreateUser(u *User, password string) (error, User)
}
