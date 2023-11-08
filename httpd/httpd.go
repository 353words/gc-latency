package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"

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

	data := make([]byte, 1<<20)

	w.Header().Set("Content-Type", "application/json")
	out := map[string]any{
		"name": user.Name,
		"id":   id,
		"size": len(data),
	}
	if json.NewEncoder(w).Encode(out); err != nil {
		log.Printf("error: can't encode user - %s", err)
	}
}

func main() {
	http.HandleFunc("/users/", userHandler)

	addr := ":8080"
	log.Printf("info: starting server with %d %s users on %s", size, db.Kind(), addr)

	if err := http.ListenAndServe(addr, nil); !errors.Is(err, http.ErrServerClosed) {
		log.Printf("error: can't run server - %s", err)
		os.Exit(1)
	}
}
