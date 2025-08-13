package http

import (
    "database/sql"
    "net/http"
)

// NewMux returns the base http.Handler with minimal routes.
func NewMux(db *sql.DB) http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        if err := db.Ping(); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            _, _ = w.Write([]byte("db: unavailable"))
            return
        }
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("ok"))
    })

    mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte(`{"pong":true}`))
    })

    // TODO: mount ingredient/dish/menu handlers here (e.g., mux.Handle("/ingredients", ...))

    return loggingMiddleware(mux)
}
