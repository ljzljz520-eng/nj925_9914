package gallery

import (
	"errors"
	"fmt"
	"privatealbum/internal/model"
	"privatealbum/internal/store"
	"strings"
)

type Service struct {
	Store *store.Store
	clock int64
}

func New(s *store.Store) *Service { return &Service{Store: s, clock: 100} }
func (s *Service) Register(owner, title, code string) (model.Record, error) {
	r := model.Record{ID: model.NewID(fmt.Sprintf("%d:%s:%s", s.clock, owner, title)), Owner: strings.TrimSpace(owner), Title: strings.TrimSpace(title), AccessCodes: []string{strings.TrimSpace(code)}, Status: model.Pending, CreatedAt: s.clock, UpdatedAt: s.clock}
	s.clock++
	if err := r.Validate(); err != nil {
		return r, err
	}
	return r, s.Store.PutRecord(r)
}
func (s *Service) Search(q string) ([]model.Record, error) {
	rs, e := s.Store.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range rs {
		if r.Status != model.Archived && store.Match(r, q) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Service) Get(id string) (model.Record, error) { return s.Store.GetRecord(id) }
func (s *Service) ChangeCode(id, code string) (model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if strings.TrimSpace(code) == "" {
		return r, errors.New("code required")
	}
	r.AccessCodes = append(r.AccessCodes, strings.TrimSpace(code))
	r.UpdatedAt = s.clock
	s.clock++
	return r, s.Store.PutRecord(r)
}
func (s *Service) Publish(id string) (map[string]string, error) {
	r, e := s.Get(id)
	if e != nil {
		return nil, e
	}
	if r.Status != model.Approved {
		return nil, errors.New("record not approved")
	}
	return r.PublicSummary(), nil
}
func (s *Service) Count() int {
	rs, e := s.Store.ListRecords()
	if e != nil {
		return 0
	}
	return len(rs)
}
