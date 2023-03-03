package server

import (
	"net/http"

	"github.com/gameloungee/server/config"
	"github.com/gameloungee/server/internal/handler"
)

type Server struct {
	base http.Server
}

func (s *Server) Run(addr string) error {
	s.base = http.Server{
		Addr:    addr,
		Handler: handler.InitHanlder(),
	}

	return s.base.ListenAndServe()
}

func (s *Server) MakeAddr() string {
	conf := config.New()

	if conf.AppMode == config.PROD_MOD {
		return ":" + conf.Port
	}

	return "localhost:" + conf.Port
}
