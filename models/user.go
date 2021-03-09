package models

// User type is used to define the structure of a user object
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers returns list of all registered users
func GetUsers() []*User {
	return users
}

// AddUser allows adding a new user
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}
