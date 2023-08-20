package gclatency

import (
	"runtime"
	"testing"

	"github.com/353words/gc-latency/users"
)

var (
	db = users.NewDB(1_000_000)
)

func BenchmarkGC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.GC()
	}
}
