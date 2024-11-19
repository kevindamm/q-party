package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"
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

	//
	//
	//
	jarchive_json, err := f.ReadFile("jarchive.json")
	if err != nil {
		log.Fatal(err)
	}
	jarchive, err := qparty.LoadJArchiveIndex(jarchive_json)
	if err != nil {
		log.Fatal(err)
	}
	jarchive.Categories = make(map[string]qparty.CategoryMetadata)
	files, err := os.ReadDir("json/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		bytes, err := os.ReadFile(path.Join("json", file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		var episode qparty.FullEpisode
		err = json.Unmarshal(bytes, &episode)
		if err != nil {
			log.Fatal(err)
		}
		year := 0
		if !episode.Aired.IsZero() {
			year = episode.Aired.Year
		}
		if year == 0 {
			continue
		}
		if episode.Single != nil {
			for _, category := range episode.Single.Columns {
				if !category.Complete() {
					continue
				}
				metadata, ok := jarchive.Categories[category.Title]
				if !ok {
					metadata = qparty.CategoryMetadata{
						Title: category.Title,
						Theme: qparty.ThemeUnknown,
						Episodes: []qparty.CategoryAired{{
							EpisodeID: episode.EpisodeID,
							ShowDate:  episode.Aired}}}
				} else {
					metadata.Episodes = append(
						metadata.Episodes,
						qparty.CategoryAired{
							EpisodeID: episode.EpisodeID,
							ShowDate:  episode.Aired},
					)
				}
				jarchive.Categories[category.Title] = metadata
			}
		} else {
			fmt.Println("episode without a Single J: ", strconv.Itoa(int(episode.EpisodeID)), episode.ShowTitle)
		}
		if episode.Double != nil {
			for _, category := range episode.Double.Columns {
				if !category.Complete() {
					continue
				}
				metadata, ok := jarchive.Categories[category.Title]
				if !ok {
					metadata = qparty.CategoryMetadata{
						Title: category.Title,
						Theme: qparty.ThemeUnknown,
						Episodes: []qparty.CategoryAired{{
							EpisodeID: episode.EpisodeID,
							ShowDate:  episode.Aired}}}
				} else {
					metadata.Episodes = append(
						metadata.Episodes,
						qparty.CategoryAired{
							EpisodeID: episode.EpisodeID,
							ShowDate:  episode.Aired},
					)
				}
				jarchive.Categories[category.Title] = metadata
			}
		} else {
			fmt.Println("episode without a Double J: ", strconv.Itoa(int(episode.EpisodeID)), episode.ShowTitle)
		}
	}
	writer, err := os.Create("jarchive2.json")
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := json.Marshal(jarchive)
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(bytes)
	writer.Close()
	//
	//
	//
	// TODO Remove
	if true {
		return
	}
	//
	//

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
var f embed.FS

func NewServer() (*service.Server, error) {
	jarchive_json, err := f.ReadFile("jarchive.json")
	if err != nil {
		return nil, err
	}
	jarchive := new(qparty.JArchiveIndex)
	err = json.Unmarshal(jarchive_json, jarchive)
	if err != nil {
		return nil, err
	}
	server := service.NewServer(jarchive)

	// TODO additional indices, set up cache, etc.

	return server, nil
}
