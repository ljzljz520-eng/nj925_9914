package store

import (
	"path/filepath"
	"privatealbum/internal/model"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "a.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.Record{ID: "persist", Owner: "owner", Title: "album", AccessCodes: []string{"code"}, Status: model.Pending}
	if e = s.PutRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetRecord("persist")
	if e != nil || got.Title != "album" {
		t.Fatalf("%v %#v", e, got)
	}
}
