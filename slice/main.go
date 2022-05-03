package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	users []string
)

func allocUsers(size int) {
	users = make([]string, size)
	for i := 0; i < size; i++ {
		users[i] = fmt.Sprintf("user-%06d", i)
	}
}

func getUser(id int) string {
	return users[id]
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

	/*
		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)
		fmt.Println("GC pause", ms.PauseTotalNs, "ns")
	*/
}
