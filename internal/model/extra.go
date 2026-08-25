package model

import (
	"sort"
	"strings"
)

func NormalizeOwner(v string) string { return strings.Title(strings.ToLower(strings.TrimSpace(v))) }
func SortRecords(rs []Record) []Record {
	out := append([]Record(nil), rs...)
	sort.Slice(out, func(i, j int) bool { return out[i].Title < out[j].Title })
	return out
}
func CloneRecord(r Record) Record { r.AccessCodes = append([]string(nil), r.AccessCodes...); return r }
