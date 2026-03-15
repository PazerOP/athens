package external

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gomods/athens/pkg/storage/mem"
	"github.com/wow-look-at-my/testify/require"
)

func TestNewServer_List(t *testing.T) {
	s, err := mem.NewStorage()
	require.NoError(t, err)

	handler := NewServer(s)
	req := httptest.NewRequest(http.MethodGet, "/github.com/test/mod/@v/list", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestNewServer_Info_NotFound(t *testing.T) {
	s, err := mem.NewStorage()
	require.NoError(t, err)

	handler := NewServer(s)
	req := httptest.NewRequest(http.MethodGet, "/github.com/test/mod/@v/v1.0.0.info", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestNewServer_GoMod_NotFound(t *testing.T) {
	s, err := mem.NewStorage()
	require.NoError(t, err)

	handler := NewServer(s)
	req := httptest.NewRequest(http.MethodGet, "/github.com/test/mod/@v/v1.0.0.mod", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestNewServer_Zip_NotFound(t *testing.T) {
	s, err := mem.NewStorage()
	require.NoError(t, err)

	handler := NewServer(s)
	req := httptest.NewRequest(http.MethodGet, "/github.com/test/mod/@v/v1.0.0.zip", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestNewServer_Delete_NotFound(t *testing.T) {
	s, err := mem.NewStorage()
	require.NoError(t, err)

	handler := NewServer(s)
	req := httptest.NewRequest(http.MethodDelete, "/github.com/test/mod/@v/v1.0.0.delete", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestNewServer_WithData(t *testing.T) {
	s, err := mem.NewStorage()
	require.NoError(t, err)

	ctx := context.Background()
	modContent := []byte("module github.com/test/mod")
	infoContent := []byte(`{"Version":"v1.0.0"}`)

	zipContent := bytes.NewReader([]byte("fake zip content"))
	err = s.Save(ctx, "github.com/test/mod", "v1.0.0", modContent, zipContent, nil, infoContent)
	require.NoError(t, err)

	handler := NewServer(s)

	// Test list
	req := httptest.NewRequest(http.MethodGet, "/github.com/test/mod/@v/list", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), "v1.0.0")

	// Test info
	req = httptest.NewRequest(http.MethodGet, "/github.com/test/mod/@v/v1.0.0.info", nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	// Test mod
	req = httptest.NewRequest(http.MethodGet, "/github.com/test/mod/@v/v1.0.0.mod", nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}
