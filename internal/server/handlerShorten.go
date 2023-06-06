package server

import (
	"net/http"
	"log"
	"github.com/errrov/urlshortenerozon/internal/model"
	"github.com/errrov/urlshortenerozon/internal/shorten"
	"github.com/labstack/echo/v4"
)

type shortener interface {
	AddShortenLinkToStorage(model.UserInput) (*model.Shortened, error)
}

type ShortenRequest struct {
	URL        string `json:"url" query:"url" param:"url" validate:"required,url"`
	Identifier string `json:"identifier,omitempty" query:"identifier" param:"identifier" validate:"omitempty,alphanum"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

func HandlerShorten(shortener shortener) echo.HandlerFunc {
	return func(c echo.Context) error {
		var shortenReq ShortenRequest
		if err := c.Bind(&shortenReq); err != nil {
			return err
		}
		log.Println("Parsed REQ:", shortenReq)
		var identifier string
		if shortenReq.Identifier != "" {
			identifier = shortenReq.Identifier
		}

		userInput := model.UserInput{
			OriginalURL: shortenReq.URL,
			Identifier:  identifier,
		}
		log.Println("IN test", userInput, identifier)
		shortening, err := shortener.AddShortenLinkToStorage(userInput)
		if err != nil {
			if err == model.ErrIdExist {
				return echo.NewHTTPError(http.StatusConflict, model.ErrIdExist.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		shortURL, err := shorten.CreateShortURL("localhost", shortening.Identifier)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(
			http.StatusOK,
			shortenResponse{ShortURL: shortURL},
		)
	}
}
