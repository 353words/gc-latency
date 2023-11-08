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
	users []User
}

func NewDB(size int) *DB {
	users := make([]User, size)
	for i := 0; i < size; i++ {
		users[i] = User{ID: i, Name: fmt.Sprintf("user-%06d", i)}
	}

	db := DB{
		users: users,
	}
	return &db
}

// ByID returns user name from id.
func (db *DB) ByID(id int) (User, bool) {
	if id < 0 || id > len(db.users) {
		return User{}, false
	}

	return db.users[id], true
}

func (*DB) Kind() string {
	return "slice"
}
