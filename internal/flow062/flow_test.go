package flow062

import (
	"path/filepath"
	"privatealbum/internal/importer"
	"privatealbum/internal/model"
	"privatealbum/internal/store"
	"testing"
)

func TestWorkflowCreateReviewArchive(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r, e := New(s).CreateReviewArchive("a", "t", "c", "mod")
	if e != nil || r.Status != model.Archived {
		t.Fatal(e, r)
	}
}
func TestWorkflowSearchUpdatePublish(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	f := New(s)
	r, _ := f.Gallery.Register("a", "t", "c")
	f.Review.Decide(r.ID, "m", "approve", "ok")
	m, e := f.SearchUpdatePublish("a", r.ID, "new")
	if e != nil || m["code"] != "new" {
		t.Fatal(e, m)
	}
}
func TestWorkflowImportReport(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := New(s).ImportReport([]importer.Row{{Owner: "a", Title: "t", Code: "c"}})
	if !r.Successful() {
		t.Fatal(r)
	}
}
