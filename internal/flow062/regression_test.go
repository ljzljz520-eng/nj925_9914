package flow062

import (
	"path/filepath"
	"privatealbum/internal/store"
	"testing"
)

func Test925BusinessRegression(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	f := New(s)
	a, _ := f.Gallery.Register("owner", "first", "first-code")
	b, _ := f.Gallery.Register("owner", "second", "second-code")
	if _, e := f.Gallery.ChangeCode(a.ID, "first-new"); e != nil {
		t.Fatal(e)
	}
	got, e := f.Gallery.ChangeCode(b.ID, "second-new")
	if e != nil {
		t.Fatal(e)
	}
	if len(got.AccessCodes) != 2 || got.CurrentCode() != "second-new" {
		t.Fatalf("unexpected access code state: %#v", got.AccessCodes)
	}
}
