package users_str

import (
	"strings"
)

var (
	users   string
	indices []int
)

func AllocateUsers(size int, userName func(int) string) {
	var buf strings.Builder
	indices = make([]int, size)
	s := 0
	for i := 0; i < size; i++ {
		indices[i] = s
		u := userName(i)
		buf.WriteString(u)
		s += len(u)
	}
	users = buf.String()
}

func ByID(id int) string {
	start := indices[id]
	if id == len(indices) {
		return users[start:]
	}
	end := indices[id+1]
	return users[start:end]
}
