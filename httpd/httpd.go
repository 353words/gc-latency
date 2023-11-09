package main

import (
	"encoding/json"
	"errors"
	"log"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"

	"github.com/353words/gc-latency/users"
)

type API struct {
	db  *users.DB
	log *slog.Logger
}

const (
	userPrefix = "/user/"
)

func (a *API) UserHandler(w http.ResponseWriter, r *http.Request) {
	n := len(userPrefix)
	if len(r.URL.Path) <= n {
		a.log.Warn("bad request", "error", "missing uid")
		http.Error(w, "missing uid", http.StatusBadRequest)
	}

	uid, err := strconv.Atoi(r.URL.Path[len(userPrefix):])
	if err != nil {
		a.log.Warn("bad request", "error", "missing uid")
		http.Error(w, "missing parameter", http.StatusBadRequest)
	}
	user, ok := a.db.ByID(uid)
	if !ok {
		a.log.Warn("unknown user", "uid", uid)
		http.Error(w, "unknown user", http.StatusNotFound)
		return
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size <= 0 {
		size = 1 << 20 // 1MB
	}
	// Simulate work generating data on heap
	data := make([]byte, size)

	w.Header().Set("Content-Type", "application/json")
	out := map[string]any{
		"name": user.Name,
		"id":   uid,
		"size": len(data),
	}
	if err := json.NewEncoder(w).Encode(out); err != nil {
		log.Printf("error: can't encode user - %s", err)
	}
}

func main() {

	const size = 1_000_000
	api := API{
		db:  users.NewDB(size),
		log: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	mux := http.NewServeMux()
	mux.HandleFunc(userPrefix, api.UserHandler)

	addr := ":8080"
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	api.log.Info("server starting", "num users", size, "kind", api.db.Kind(), "addr", addr)
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Printf("error: can't run server - %s", err)
		os.Exit(1)
	}
}
