package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	server := NewServer(port)

	if err := server.Run(); err != nil {
		return
	}
}

type Server struct {
	port string
}

func NewServer(port string) *Server {
	s := &Server{
		port: port,
	}

	return s
}

func (s *Server) Run() error {
	http.HandleFunc("/", s.getRoot)

	return http.ListenAndServe(":"+s.port, nil)
}

func (s *Server) getRoot(writer http.ResponseWriter, requst *http.Request) {
	header := writer.Header()
	header.Set("Content-Type", "text/plain")

	fmt.Fprintln(writer, "hello, world!")
}
