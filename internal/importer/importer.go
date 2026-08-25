package importer

import (
	"fmt"
	"privatealbum/internal/gallery"
	"privatealbum/internal/store"
	"strings"
)

type Row struct{ Owner, Title, Code string }
type Result struct {
	Accepted int
	Rejected int
	Errors   []string
	IDs      []string
}
type Service struct {
	Gallery *gallery.Service
	Store   *store.Store
}

func New(g *gallery.Service, s *store.Store) *Service { return &Service{Gallery: g, Store: s} }
func (s *Service) ImportRows(rows []Row) Result {
	res := Result{}
	for i, row := range rows {
		if strings.TrimSpace(row.Owner) == "" || strings.TrimSpace(row.Title) == "" || strings.TrimSpace(row.Code) == "" {
			res.Rejected++
			res.Errors = append(res.Errors, fmt.Sprintf("row %d: required fields", i))
			continue
		}
		r, e := s.Gallery.Register(row.Owner, row.Title, row.Code)
		if e != nil {
			res.Rejected++
			res.Errors = append(res.Errors, fmt.Sprintf("row %d: %v", i, e))
			continue
		}
		res.Accepted++
		res.IDs = append(res.IDs, r.ID)
	}
	return res
}
func ValidateRows(rows []Row) []string {
	errs := []string{}
	for i, r := range rows {
		if strings.TrimSpace(r.Owner) == "" {
			errs = append(errs, fmt.Sprintf("%d owner", i))
		}
		if strings.TrimSpace(r.Title) == "" {
			errs = append(errs, fmt.Sprintf("%d title", i))
		}
		if strings.TrimSpace(r.Code) == "" {
			errs = append(errs, fmt.Sprintf("%d code", i))
		}
	}
	return errs
}
