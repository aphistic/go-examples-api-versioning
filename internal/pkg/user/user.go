package user

type User struct {
	ID    string
	Email string
	Name  string
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(user *User) error {
	// Any code we need to create users
	return nil
}

func (us *UserService) ListUsers() ([]*User, error) {
	// Get all our users, this would probably have some kind of filter or paging on it
	return []*User{
		{ID: "id", Email: "email", Name: "name"},
	}, nil
}

func (us *UserService) GetUserByID(id string) (*User, error) {
	// Get the user by the ID, probably not a string but that's easiest
	// for this example because it's not really important
	return &User{ID: "id", Email: "email", Name: "name"}, nil
}
