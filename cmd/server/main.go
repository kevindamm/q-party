package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	qparty "github.com/kevindamm/q-party"
	"github.com/kevindamm/q-party/service"
)

// Calls the server's [Shutdown()] when
// Runs in a goroutine alongside the server handler, the [done] channel runs the
// remainder of main() (after blocking on a channel following server startup).
// This gives a convenient place to
func graceful_shutdown(https_server *http.Server, done chan<- bool, timeout_seconds int) {
	// Listen for the interrupt signal or termination from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// Notify the server with a 5 second timeout so that current handlers can finish.
	timeout := time.Duration(timeout_seconds) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	log.Println("Graceful Shutdown: server exiting in", timeout_seconds, "...")
	if err := https_server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	// Shutdown is complete, safely notify the code execution of main().
	done <- true
}

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}

	jarchive_json, err := embedded_files.ReadFile("jarchive.json")
	if err != nil {
		log.Fatal(err)
	}
	jarchive, err := qparty.LoadJArchiveIndex(jarchive_json)
	if err != nil {
		log.Fatal(err)
	}
	server.JArchiveIndex = jarchive

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine that exits after timeout.
	go graceful_shutdown(server.Server, done, 5)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}

//go:embed episodes/*
//go:embed jarchive.json
//go:embed index.html
//go:embed favicon.ico
var embedded_files embed.FS

func NewServer() (*service.Server, error) {
	jarchive_json, err := embedded_files.ReadFile("jarchive.json")
	if err != nil {
		return nil, err
	}
	jarchive := new(qparty.JArchiveIndex)
	err = json.Unmarshal(jarchive_json, jarchive)
	if err != nil {
		return nil, err
	}
	server := service.NewServer(jarchive, embedded_files)

	// TODO additional indices, set up cache, etc.

	return server, nil
}
