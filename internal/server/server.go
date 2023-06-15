package server

import (
	"net/http"

	"github.com/errrov/urlshortener/internal/shorten"
	"github.com/labstack/echo/v4"
)

type Server struct {
	shortening *shorten.Service
	E *echo.Echo
}

func New(shortener *shorten.Service) *Server {
	s := &Server{
		shortening: shortener,
	}
	s.SetupRoutes()
	return s
}

func (s *Server) SetupRoutes() {
	s.E = echo.New()
	s.E.POST("/", HandlerShorten(s.shortening))
	s.E.GET("/", HandlerRedirect(s.shortening))

}

func (s *Server) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	s.E.ServeHTTP(w,r)
}


