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

	//middleware?
	
	
	//Routes
	s.E.POST("/", HandlerShorten(s.shortening))
	s.E.GET("/", HandlerRedirect(s.shortening))

}

func (s *Server) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	s.E.ServeHTTP(w,r)
}


