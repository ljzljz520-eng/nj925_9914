package flow062

import (
	"fmt"
	"privatealbum/internal/model"
)

func FormatSummary(r model.Record) string {
	return fmt.Sprintf("%s/%s [%s] code=%s", r.Owner, r.Title, r.Status, r.CurrentCode())
}
func StatusLabel(status string) string {
	switch status {
	case model.Pending:
		return "awaiting review"
	case model.Approved:
		return "approved"
	case model.Rejected:
		return "rejected"
	case model.Archived:
		return "archived"
	default:
		return "unknown"
	}
}
