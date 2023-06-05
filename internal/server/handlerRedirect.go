package server

import (
	"net/http"

	"github.com/errrov/urlshortenerozon/internal/model"
	"github.com/errrov/urlshortenerozon/internal/shorten"
	"github.com/labstack/echo/v4"
)

func HandlerRedirect(s *shorten.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		identifier := c.Param("identifier")
		shortened, err := s.Storage.Get(identifier)
		if err != nil {
			if err == model.ErrNotFound {
				return echo.NewHTTPError(http.StatusNotFound)
			}
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusMovedPermanently, shortened.Original)
	}
}
