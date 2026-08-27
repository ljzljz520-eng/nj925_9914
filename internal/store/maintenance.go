package store

import (
	"privatealbum/internal/model"
	"sort"
)

func SortByUpdated(rs []model.Record) []model.Record {
	out := append([]model.Record(nil), rs...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].UpdatedAt > out[j].UpdatedAt })
	return out
}
func GroupByStatus(rs []model.Record) map[string][]model.Record {
	out := map[string][]model.Record{}
	for _, r := range rs {
		out[r.Status] = append(out[r.Status], r)
	}
	return out
}
func Latest(rs []model.Record) model.Record {
	if len(rs) == 0 {
		return model.Record{}
	}
	out := SortByUpdated(rs)
	return out[0]
}
