package users

import "fmt"

const (
	nameFmt = "user-%06d"
)

type User struct {
	ID   int
	Name string
}

type DB struct {
	users map[int]User
}

// NewDB create a new UserDB with "size" users.
func NewDB(size int) *DB {
	users := make(map[int]User)
	for i := 0; i < size; i++ {
		users[i] = User{ID: i, Name: fmt.Sprintf(nameFmt, i)}
	}

	db := DB{
		users: users,
	}
	return &db
}

// ByID returns user name from id.
func (db *DB) ByID(id int) (User, bool) {
	user, ok := db.users[id]
	return user, ok
}
