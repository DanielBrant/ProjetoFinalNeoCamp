package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    ih "ProjetoFinalNeoCamp/internal/http"
    storemysql "ProjetoFinalNeoCamp/internal/store/mysql"
)

var (
    version = "dev"
    commit  = "unknown"
)

func main() {
    addr := getEnv("HTTP_ADDR", ":8080")

    // DB setup
    db, err := storemysql.OpenFromEnv()
    if err != nil {
        log.Fatalf("open db: %v", err)
    }
    defer db.Close()

    // Router
    mux := ih.NewMux(db)

    srv := &http.Server{
        Addr:              addr,
        Handler:           mux,
        ReadHeaderTimeout: 5 * time.Second,
    }

    go func() {
        log.Printf("ProjetoFinalNeoCamp %s (%s) listening on %s", version, commit, addr)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("server: %v", err)
        }
    }()

    // Graceful shutdown
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    <-stop
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Printf("shutdown: %v", err)
    }
    log.Println("server stopped")
}

func getEnv(k, def string) string {
    if v := os.Getenv(k); v != "" {
        return v
    }
    return def
}
