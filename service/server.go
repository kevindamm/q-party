package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	qparty "github.com/kevindamm/q-party"
)

type Server struct {
	port int
}

func NewServer(jarchive *qparty.JArchiveIndex) *http.Server {
	port, err := strconv.Atoi(os.Getenv("QPARTY_PORT"))
	if err != nil {
		port = 0
	}
	QPartyServer := &Server{port: port}

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", QPartyServer.port),
		Handler:      QPartyServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return &server
}
