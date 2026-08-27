package review

import (
	"privatealbum/internal/model"
	"strings"
)

func CanDecide(actor string, r model.Record) bool {
	return strings.TrimSpace(actor) != "" && r.Status == model.Pending
}
func DecisionLabel(v string) string {
	switch v {
	case "approve":
		return "Approved"
	case "reject":
		return "Rejected"
	case "defer":
		return "Deferred"
	default:
		return "Unknown"
	}
}
