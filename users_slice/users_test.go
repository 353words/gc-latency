package users

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestByID(t *testing.T) {
	size := 100
	db := NewDB(size)
	id := 7
	user, ok := db.ByID(id)
	require.Truef(t, ok, "get user %d", id)
	expected := User{ID: id, Name: fmt.Sprintf(nameFmt, id)}
	require.Equal(t, expected, user, "user name")

	_, ok = db.ByID(size + 3)
	require.False(t, ok, "non existing user")
}

func BenchmarkByID(b *testing.B) {
	db := NewDB(100)
	id := 42
	for i := 0; i < b.N; i++ {
		u, ok := db.ByID(id)
		if !ok || u.ID != id {
			b.Fatal(u, ok)
		}
	}
}
