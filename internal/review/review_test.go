package review

import (
	"path/filepath"
	"privatealbum/internal/gallery"
	"privatealbum/internal/model"
	"privatealbum/internal/store"
	"testing"
)

func TestDecisionAudit(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r, _ := gallery.New(s).Register("a", "t", "c")
	got, e := New(s).Decide(r.ID, "mod", "approve", "ok")
	if e != nil || got.Status != model.Approved {
		t.Fatal(e, got)
	}
	hs, _ := New(s).History(r.ID)
	if len(hs) != 1 {
		t.Fatal("audit missing")
	}
}
