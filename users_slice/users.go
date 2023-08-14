package users

var (
	users []string
)

func AllocateUsers(size int, userName func(int) string) {
	users = make([]string, size)
	for i := 0; i < size; i++ {
		users[i] = userName(i)
	}
}

func ByID(id int) string {
	return users[id]
}
