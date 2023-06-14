package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/errrov/urlshortener/internal/model"
	"github.com/errrov/urlshortener/internal/server"
	"github.com/errrov/urlshortener/internal/shorten"
	"github.com/errrov/urlshortener/internal/storage/in_memory"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerShortenBase(t *testing.T) {
	payload := `{"url":"https://www.google.com","identifier":"google"}`
	shortener := shorten.NewService(in_memory.NewInMemory())
	recorder := httptest.NewRecorder()
	e := echo.New()

	handler := server.HandlerShorten(shortener)
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	var response struct {
		ShortURL string `json:"short_url"`
	}
	c := e.NewContext(request, recorder)
	require.NoError(t, handler(c))

	assert.Equal(t, http.StatusOK, recorder.Code)
	require.NoError(t, json.NewDecoder(recorder.Body).Decode(&response))
	assert.NotEmpty(t, response.ShortURL)
}

func TestHandlerShortenIfIdTaken(t *testing.T) {
	payload := `{"url":"https://www.google.com","identifier":"google"}`
	shortener := shorten.NewService(in_memory.NewInMemory())
	recorder := httptest.NewRecorder()
	e := echo.New()

	handler := server.HandlerShorten(shortener)
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(request, recorder)
	require.NoError(t, handler(c))
	assert.Equal(t, http.StatusOK, recorder.Code)
	request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder = httptest.NewRecorder()
	c = e.NewContext(request, recorder)
	var httpError *echo.HTTPError
	require.ErrorAs(t, handler(c), &httpError)
	assert.Equal(t, http.StatusConflict, httpError.Code)
	assert.Contains(t, httpError.Message, model.ErrIdExist.Error())
}
