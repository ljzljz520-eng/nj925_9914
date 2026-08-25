package store

import (
	"privatealbum/internal/model"
	"strings"
)

func FilterStatus(rs []model.Record, status string) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if status == "" || r.Status == status {
			out = append(out, r)
		}
	}
	return out
}
func SearchByOwner(rs []model.Record, owner string) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if strings.EqualFold(r.Owner, owner) {
			out = append(out, r)
		}
	}
	return out
}
