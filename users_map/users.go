package users

var (
	users map[int]string
)

func AllocateUsers(size int, userName func(int) string) {
	users = make(map[int]string)
	for i := 0; i < size; i++ {
		users[i] = userName(i)
	}
}

func ByID(id int) string {
	return users[id]
}
