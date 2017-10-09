package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/olefine/go-dockerized-web/handlers"
)

// Server instance to run webserver from outline
type Server struct {
	Port uint16
	Host string
}

func (server Server) String() string {
	return fmt.Sprintf(":%d", server.Port)
}

// NewServer creates new server instance, host and port should be provided
func NewServer(host string, port uint16) *Server {
	return &Server{Host: host, Port: port}
}

// Run starts event loop for Server
func (server Server) Run() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/persons", handlers.Persons)

	log.Println(http.ListenAndServe(server.String(), serverMux))
}
