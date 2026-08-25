package review

import (
	"errors"
	"fmt"
	"privatealbum/internal/model"
	"privatealbum/internal/store"
)

type Service struct {
	Store *store.Store
	now   int64
}

func New(s *store.Store) *Service { return &Service{Store: s, now: 500} }
func (s *Service) Decide(id, actor, decision, detail string) (model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if r.Status != model.Pending {
		return r, errors.New("record is not pending")
	}
	switch decision {
	case "approve":
		r.Status = model.Approved
	case "reject":
		r.Status = model.Rejected
	case "defer":
		return r, nil
	default:
		return r, fmt.Errorf("unknown decision %q", decision)
	}
	r.UpdatedAt = s.now
	s.now++
	if e = s.Store.PutRecord(r); e != nil {
		return r, e
	}
	a := model.AuditEvent{ID: model.NewID(fmt.Sprintf("%s:%d", id, s.now)), RecordID: id, Actor: actor, Action: decision, Detail: detail, At: s.now}
	return r, s.Store.PutAudit(a)
}
func (s *Service) History(id string) ([]model.AuditEvent, error) { return s.Store.ListAudit(id) }
