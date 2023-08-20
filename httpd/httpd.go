package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/353words/gc-latency/users"
)

const size = 1_000_000

var (
	db = users.NewDB(size)
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	const prefix = "/users/"
	id, err := strconv.Atoi(r.URL.Path[len(prefix):])
	if err != nil {
		http.Error(w, "bad user ID", http.StatusBadRequest)
		return
	}

	user, ok := db.ByID(id)
	if !ok {
		http.Error(w, "unknown user", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	out := map[string]any{
		"name": user.Name,
		"id":   id,
	}
	if json.NewEncoder(w).Encode(out); err != nil {
		log.Printf("error: can't encode user - %s", err)
	}
}

/* Generate random data so it'll trigger GC
func spammer() {
	var s []float64
	for {
		time.Sleep(time.Millisecond)

		s = make([]float64, rand.Intn(100_000))
		for i := 0; i < len(s); i++ {
			s[i] = rand.Float64()
		}
	}
}
*/

// Trigger GC from time to time
func GCer() {
	for {
		time.Sleep(100 * time.Millisecond)
		runtime.GC()
	}
}

func main() {
	http.HandleFunc("/users/", userHandler)

	go GCer()

	addr := ":8080"
	log.Printf("info: starting server with %d users on %s", size, addr)

	if err := http.ListenAndServe(addr, nil); !errors.Is(err, http.ErrServerClosed) {
		log.Printf("error: can't run server - %s", err)
		os.Exit(1)
	}
}
