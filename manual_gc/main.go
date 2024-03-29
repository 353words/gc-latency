package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"time"

	"github.com/353words/gc-latency/users"
)

const size = 1_000_000

var (
	db = users.NewDB(size)
)

func main() {
	id := 353
	u, ok := db.ByID(id)
	if !ok || u.ID != id {
		log.Fatalf("error: bad user %#v (%v)", u, ok)
	}
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
	// avg := time.Duration(float64(total) / float64(len(durations)))
	fmt.Printf("total gc time: %v (median: %v)\n", total, median(durations))

	/*
		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)
		fmt.Println("GC pause", ms.PauseTotalNs, "ns")
	*/
}

func median(ds []time.Duration) time.Duration {
	cp := make([]time.Duration, len(ds))
	copy(cp, ds)
	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })

	i := len(cp) / 2
	if len(cp)%2 == 1 {
		return cp[i]
	}

	return (cp[i] + cp[i+1]) / 2
}

// 4MB of memory at least (so GC will trigger)
// Run through the tracer
// Look at pace, not GC time
// Turn off GC and see effect
// Percent time in GC
// https://github.com/ardanlabs/gotraining/blob/master/topics/go/profiling/trace/trace.go
