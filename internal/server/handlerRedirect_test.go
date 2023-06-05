package server_test

import (
	"net/http"
	"testing"

	"github.com/errrov/urlshortenerozon/internal/model"
	"github.com/errrov/urlshortenerozon/internal/server"
	"github.com/errrov/urlshortenerozon/internal/shorten"
	"github.com/errrov/urlshortenerozon/internal/storage/in_memory"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
)

func TestHandlerRedirect(t *testing.T) {
	identifier, url := "google", "https://www.google.com" 
	redirect := shorten.NewService(in_memory.NewInMemory())
	recorder := httptest.NewRecorder()
	e := echo.New()

	handler := server.HandlerRedirect(redirect)
	request := httptest.NewRequest(http.MethodGet, "/"+identifier, nil)
	
	c := e.NewContext(request, recorder)
	c.SetPath("/:identifier")
	c.SetParamNames("identifier")
	c.SetParamValues(identifier)
	_, err := redirect.AddShortenLinkToStorage(model.UserInput{
		Identifier: identifier,
		OriginalURL: url,
	},)
	require.NoError(t, err)
	require.NoError(t, handler(c))
	assert.Equal(t, http.StatusMovedPermanently, recorder.Code)
	assert.Equal(t, url, recorder.Header().Get("Location"))
}