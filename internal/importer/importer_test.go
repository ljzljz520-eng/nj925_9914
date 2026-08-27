package importer

import (
	"path/filepath"
	"privatealbum/internal/gallery"
	"privatealbum/internal/store"
	"testing"
)

func TestImportValidation(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := New(gallery.New(s), s).ImportRows([]Row{{Owner: "a", Title: "t", Code: "c"}, {Owner: "", Title: "bad", Code: "x"}})
	if r.Accepted != 1 || r.Rejected != 1 {
		t.Fatal(r)
	}
	if len(ValidateRows([]Row{{Title: "x"}})) != 2 {
		t.Fatal("validation")
	}
}
