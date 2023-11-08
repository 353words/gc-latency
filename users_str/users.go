package users

import (
	"fmt"
	"strings"
)

const (
	nameFmt = "user-%06d"
)

type User struct {
	ID   int
	Name string
}

type DB struct {
	users   string
	indices []int
}

func NewDB(size int) *DB {
	var buf strings.Builder
	indices := make([]int, size)
	s := 0
	for i := 0; i < size; i++ {
		indices[i] = s
		u := fmt.Sprintf("user-%06d", i)
		buf.WriteString(u)
		s += len(u)
	}

	db := DB{
		indices: indices,
		users:   buf.String(),
	}
	return &db
}

// ByID returns user name from id.
func (db *DB) ByID(id int) (User, bool) {
	if id < 0 || id > len(db.indices) {
		return User{}, false
	}

	start := db.indices[id]
	var name string
	if id == len(db.indices) {
		name = db.users[start:]
	} else {
		end := db.indices[id+1]
		name = db.users[start:end]
	}

	return User{ID: id, Name: name}, true
}

func (*DB) Kind() string {
	return "string"
}
