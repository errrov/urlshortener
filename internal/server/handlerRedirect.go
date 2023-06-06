package server

import (
	"log"
	"net/http"

	"github.com/errrov/urlshortenerozon/internal/model"
	"github.com/errrov/urlshortenerozon/internal/shorten"
	"github.com/labstack/echo/v4"
)

type redirectReq struct {
	Identifier string `json:"identifier" param:"identifier" query:"identifier"`
}

func HandlerRedirect(s *shorten.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var rReq redirectReq
		if err := c.Bind(&rReq); err != nil {
			return err
		}
		log.Println("Redirect Got identifier:", rReq.Identifier)
		shortened, err := s.Storage.Get(rReq.Identifier)
		if err != nil {
			if err == model.ErrNotFound {
				return echo.NewHTTPError(http.StatusNotFound)
			}
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusMovedPermanently, shortened.Original)
	}
}
