package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lzaun/iam/config"
)

type Server struct {
	server *http.Server
}

func NewServer(conf *config.Config) *Server {
	return &Server{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", conf.Port),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) Serve() {
	if err := s.server.ListenAndServe(); err != nil {
		panic(err)
	}
}
