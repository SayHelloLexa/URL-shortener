package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sayhellolexa/url-short/internal/service"
)

type Server struct {
	router  *mux.Router
	service *service.UrlService
}

func NewServer() *Server {
	s := &Server{
		router:  mux.NewRouter(),
		service: service.NewUrlService(),
	}

	s.configureRoutes()
	return s
}

func (s *Server) Start(addr string) error {
	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		return err
	}
	return nil
}
