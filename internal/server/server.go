package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	password string
}

func New(password string) *Server {
	return &Server{
		password: password,
	}
}

func (s *Server) Listen(port int) {
	log.Infof("Password: %s", s.password)
	addr := fmt.Sprintf(":%d", port)
	log.Infof("Started listening on %s\n", addr)
}
