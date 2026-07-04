package server

import (
	"fmt"
	"net/http"

	"github.com/chanelog/sshws-core/internal/config"
	"github.com/chanelog/sshws-core/internal/logger"
)

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) routes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "SSHWS Core Running")

	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "OK")

	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "SSHWS Core v1.0.0-dev")

	})

}

func (s *Server) Start() error {

	s.routes()

	logger.Info.Println("===================================")
	logger.Info.Println(" SSHWS CORE")
	logger.Info.Println("===================================")
	logger.Info.Printf("Listen : %s\n", s.cfg.Listen)

	return http.ListenAndServe(
		s.cfg.Listen,
		nil,
	)

}
