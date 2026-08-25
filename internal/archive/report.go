package archive

import (
	"privatealbum/internal/model"
	"privatealbum/internal/store"
)

func ArchivedRecords(s *store.Store) ([]model.Record, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range rs {
		if r.Status == model.Archived {
			out = append(out, r)
		}
	}
	return out, nil
}
func ActiveRecords(s *store.Store) ([]model.Record, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range rs {
		if r.Status != model.Archived {
			out = append(out, r)
		}
	}
	return out, nil
}
