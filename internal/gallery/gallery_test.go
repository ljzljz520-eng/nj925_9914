package gallery

import (
	"path/filepath"
	"privatealbum/internal/store"
	"testing"
)

func TestRegisterSearch(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	g := New(s)
	r, e := g.Register("Ana", "Summer", "one")
	if e != nil {
		t.Fatal(e)
	}
	rs, e := g.Search("summer")
	if e != nil || len(rs) != 1 || rs[0].ID != r.ID {
		t.Fatalf("%v %v", e, rs)
	}
}
