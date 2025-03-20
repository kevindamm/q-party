package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/kevindamm/q-party/service"
)

// Main entry point for the web server, loads some static data before starting.
// Wraps the server in a signal listener that initiates a graceful shutdown.

func main() {
	initialize := flag.Bool("init", false, "initialize DB from embedded metadata")
	cert_path := flag.String("cert_path", "",
		"path where server.crt and server.key can be found (for this environment)")
	port := flag.Int("port", 0,
		"server port to listen on (0 uses default 80/443 ports)")
	flag.Parse()

	//jarchive := ent.NewClient()
	if *initialize {
		//	err := jarchive.Schema.Create(context.Background())
		//	if err != nil {
		//		log.Fatal(err)
		//	}
	}

	if *port == 0 {
		*port = 80
		if len(*cert_path) > 0 {
			*port = 443
		}
	}

	server := service.NewServer(*port)
	server.LoadJArchiveIndex(must_read_file(staticFS, "static/jarchive.json"))

	server.RouteStaticFiles(staticFS)
	//server.RoutePlayRooms()
	//server.RouteMediaFiles()

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)
	// Run graceful shutdown in a separate goroutine that exits after timeout.
	go graceful_shutdown(server.Server, done, 5)

	var err error
	if len(*cert_path) > 0 {
		crt_path := path.Join(*cert_path, "server.crt")
		key_path := path.Join(*cert_path, "server.key")
		err = server.ServeTLS(crt_path, key_path)
	} else {
		// Serve without HTTPS if the certificates path is empty.
		err = server.Serve(*port)
	}
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("httpserver error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("graceful shutdown complete")
}

//go:embed static/*
var staticFS embed.FS

// Calls the server's [Shutdown()] when a SIGINT or SIGTERM signal is caught.
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

// Reads the named file from the indicated filesystem.
// (works for both embedded and system-local files)
func must_read_file(fs fs.FS, filename string) []byte {
	reader, err := fs.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
