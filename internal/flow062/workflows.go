package flow062

import (
	"errors"
	"privatealbum/internal/importer"
	"privatealbum/internal/model"
)

func (s *Service) CreateReviewArchive(owner, title, code, actor string) (model.Record, error) {
	r, e := s.Gallery.Register(owner, title, code)
	if e != nil {
		return r, e
	}
	r, e = s.Review.Decide(r.ID, actor, "approve", "initial review")
	if e != nil {
		return r, e
	}
	return s.Archive.Archive(r.ID)
}
func (s *Service) SearchUpdatePublish(q, id, code string) (map[string]string, error) {
	rs, e := s.Gallery.Search(q)
	if e != nil {
		return nil, e
	}
	found := false
	for _, r := range rs {
		if r.ID == id {
			found = true
		}
	}
	if !found {
		return nil, errors.New("record not selected")
	}
	if _, e = s.Gallery.ChangeCode(id, code); e != nil {
		return nil, e
	}
	return s.Gallery.Publish(id)
}
func (s *Service) ImportReport(rows []importer.Row) importer.Result {
	return importer.New(s.Gallery, s.Gallery.Store).ImportRows(rows)
}
func (s *Service) ActiveCount() int { return s.Gallery.Count() }
