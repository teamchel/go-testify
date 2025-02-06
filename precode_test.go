package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandler(t *testing.T) {
	t.Run("Valid request returns 200 and non-empty body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=2", nil)
		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)

		handler.ServeHTTP(responseRecorder, req)

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		assert.NotEmpty(t, responseRecorder.Body.String())
	})

	t.Run("Unsupported city returns 400 and error message", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/cafe?city=unknown&count=2", nil)
		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)

		handler.ServeHTTP(responseRecorder, req)

		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
		assert.Equal(t, "wrong city value", responseRecorder.Body.String())
	})

	t.Run("Count greater than total returns all cafes", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=10", nil)
		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)

		handler.ServeHTTP(responseRecorder, req)

		requiredCafes := cafeList["moscow"]
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		require.Len(t, strings.Split(responseRecorder.Body.String(), ","), len(requiredCafes))
	})
}
