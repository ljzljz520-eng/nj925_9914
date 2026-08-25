package main

import (
	"net/http/httptest"
	"path/filepath"
	"privatealbum/internal/flow062"
	"privatealbum/internal/store"
	"testing"
)

func TestHTTPRoutes(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	rr := httptest.NewRecorder()
	flow062.New(s).Handler().ServeHTTP(rr, httptest.NewRequest("GET", "/health", nil))
	if rr.Code != 200 {
		t.Fatal(rr.Code)
	}
}
