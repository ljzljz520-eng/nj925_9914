package archive

import (
	"path/filepath"
	"privatealbum/internal/gallery"
	"privatealbum/internal/model"
	"privatealbum/internal/store"
	"testing"
)

func TestArchiveRestore(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r, _ := gallery.New(s).Register("a", "t", "c")
	r.Status = model.Approved
	s.PutRecord(r)
	a := New(s)
	a.Archive(r.ID)
	if !a.IsArchived(r.ID) {
		t.Fatal("not archived")
	}
	a.Restore(r.ID)
	if a.IsArchived(r.ID) {
		t.Fatal("not restored")
	}
}
