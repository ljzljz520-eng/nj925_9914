package gallery

import (
	"privatealbum/internal/model"
	"strings"
)

func MaskCode(code string) string {
	if len(code) <= 2 {
		return strings.Repeat("*", len(code))
	}
	return code[:1] + strings.Repeat("*", len(code)-2) + code[len(code)-1:]
}
func Visible(r model.Record) map[string]string {
	m := r.PublicSummary()
	m["code"] = MaskCode(r.CurrentCode())
	return m
}
