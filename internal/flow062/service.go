package flow062

import (
	"encoding/json"
	"net/http"
	"privatealbum/internal/archive"
	"privatealbum/internal/gallery"
	"privatealbum/internal/review"
	"privatealbum/internal/store"
	"strings"
)

type Service struct {
	Gallery *gallery.Service
	Review  *review.Service
	Archive *archive.Service
}

func New(s *store.Store) *Service {
	return &Service{Gallery: gallery.New(s), Review: review.New(s), Archive: archive.New(s)}
}
func (s *Service) Handler() http.Handler { return http.HandlerFunc(s.route) }
func (s *Service) route(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if r.Method == http.MethodGet && len(parts) == 1 && parts[0] == "health" {
		s.write(w, 200, map[string]string{"status": "ok"})
		return
	}
	if r.Method == http.MethodGet && len(parts) == 1 && parts[0] == "records" {
		rs, e := s.Gallery.Search(r.URL.Query().Get("q"))
		if e != nil {
			s.write(w, 500, map[string]string{"error": e.Error()})
			return
		}
		s.write(w, 200, rs)
		return
	}
	if r.Method == http.MethodPost && len(parts) == 1 && parts[0] == "records" {
		var in struct{ Owner, Title, Code string }
		if json.NewDecoder(r.Body).Decode(&in) != nil {
			s.write(w, 400, map[string]string{"error": "invalid json"})
			return
		}
		rec, e := s.Gallery.Register(in.Owner, in.Title, in.Code)
		if e != nil {
			s.write(w, 422, map[string]string{"error": e.Error()})
			return
		}
		s.write(w, 201, rec)
		return
	}
	s.write(w, 404, map[string]string{"error": "not found"})
}
func (s *Service) write(w http.ResponseWriter, status int, v any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
