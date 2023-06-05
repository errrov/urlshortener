package server

import (
	//"github.com/errrov/urlshortenerozon/internal/model"
	"net/http"

	"github.com/errrov/urlshortenerozon/internal/shorten"
	//"github.com/errrov/urlshortenerozon/internal/storage/in_memory"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	shortening *shorten.Service
	e *echo.Echo
}

func New(shortener *shorten.Service) *Server {
	s := &Server{
		shortening: shortener,
	}
	s.SetupRoutes()
	return s
}

func (s *Server) SetupRoutes() {
	s.e = echo.New()

	//middleware?
	
	
	//Routes
	s.e.POST("/", HandlerShorten(s.shortening))
	s.e.GET("/", HandlerRedirect(s.shortening))

}

func (s *Server) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	s.e.ServeHTTP(w,r)
}


