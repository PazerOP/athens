package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/wow-look-at-my/testify/require"
)

func TestContentType(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {}
	expected := "application/json"
	handler := ContentType(http.HandlerFunc(h))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/test", nil)
	handler.ServeHTTP(w, r)

	given := w.Result().Header.Get("Content-Type")
	require.Equal(t, expected, given)

}
