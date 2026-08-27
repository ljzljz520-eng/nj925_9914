package archive

import (
	"errors"
	"privatealbum/internal/model"
	"privatealbum/internal/store"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Archive(id string) (model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if r.Status == model.Archived {
		return r, errors.New("already archived")
	}
	r.Status = model.Archived
	r.UpdatedAt++
	return r, s.Store.PutRecord(r)
}
func (s *Service) Restore(id string) (model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if r.Status != model.Archived {
		return r, errors.New("record is active")
	}
	r.Status = model.Approved
	r.UpdatedAt++
	return r, s.Store.PutRecord(r)
}
func (s *Service) IsArchived(id string) bool {
	r, e := s.Store.GetRecord(id)
	return e == nil && r.Status == model.Archived
}
