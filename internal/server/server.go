package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lantonster/iam/config"
)

type Server struct {
	server *http.Server
}

func NewServer(conf *config.Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", conf.Port),
			Handler:      handler,
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
