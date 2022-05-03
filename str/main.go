package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

var (
	users   string
	indices []int
)

func allocUsers(size int) {
	var buf strings.Builder
	indices = make([]int, size)
	s := 0
	for i := 0; i < size; i++ {
		u := fmt.Sprintf("user-%06d", i)
		buf.WriteString(u)
		s += len(u)
		indices[i] = s
	}
	users = buf.String()
}

func getUser(id int) string {
	start, end := indices[id], indices[id+1]
	return users[start:end]
}

func main() {
	const size = 1_000_000
	allocUsers(size)
	fmt.Printf("allocated %d users\n", size)

	const nGC = 100
	fmt.Printf("running %d GC cycles\n", nGC)
	var durations []time.Duration
	for i := 0; i < nGC; i++ {
		start := time.Now()
		runtime.GC()
		duration := time.Since(start)
		durations = append(durations, duration)
	}

	var total time.Duration
	for _, d := range durations {
		total += d
	}
	fmt.Println(total)
}
