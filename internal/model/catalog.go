package model

type Rule struct {
	Code      string
	Label     string
	Required  bool
	MaxLength int
}

var AccessCodeRules = []Rule{
	{"A01", "owner supplied", true, 64}, {"A02", "title supplied", true, 120}, {"A03", "code supplied", true, 128},
	{"A04", "code printable", true, 128}, {"A05", "owner normalized", false, 80}, {"A06", "title normalized", false, 160},
	{"A07", "review actor present", true, 80}, {"A08", "review reason present", false, 500}, {"A09", "archive reason", false, 500},
	{"A10", "import row number", true, 12}, {"A11", "attachment name", true, 200}, {"A12", "attachment checksum", true, 128},
	{"A13", "workflow kind", true, 48}, {"A14", "workflow state", true, 48}, {"A15", "record identifier", true, 64},
	{"A16", "audit identifier", true, 64}, {"A17", "created timestamp", true, 20}, {"A18", "updated timestamp", true, 20},
	{"A19", "status transition", true, 32}, {"A20", "search query", false, 120}, {"A21", "page size", false, 4},
	{"A22", "page offset", false, 8}, {"A23", "sort key", false, 40}, {"A24", "sort direction", false, 8},
	{"A25", "publish channel", false, 32}, {"A26", "moderator note", false, 500}, {"A27", "restore reason", false, 500},
	{"A28", "retention class", false, 32}, {"A29", "owner contact", false, 200}, {"A30", "consent flag", true, 5},
	{"A31", "visibility flag", true, 5}, {"A32", "revision number", true, 12}, {"A33", "import batch", true, 64},
	{"A34", "source label", false, 100}, {"A35", "checksum algorithm", false, 20}, {"A36", "attachment size", true, 20},
	{"A37", "operator id", true, 80}, {"A38", "request id", true, 80}, {"A39", "locale", false, 12},
	{"A40", "timezone", false, 40}, {"A41", "error code", false, 32}, {"A42", "error message", false, 500},
	{"A43", "retry count", false, 8}, {"A44", "completion flag", true, 5}, {"A45", "approval scope", false, 80},
	{"A46", "approval version", false, 16}, {"A47", "change ticket", false, 80}, {"A48", "change reason", true, 500},
	{"A49", "previous code", false, 128}, {"A50", "current code", true, 128}, {"A51", "history count", true, 8},
	{"A52", "export format", false, 16}, {"A53", "export label", false, 100}, {"A54", "notification flag", false, 5},
	{"A55", "verification state", true, 32}, {"A56", "verification actor", false, 80}, {"A57", "verification time", false, 20},
	{"A58", "policy version", true, 16}, {"A59", "schema version", true, 8}, {"A60", "migration marker", false, 32},
}

func RuleByCode(code string) (Rule, bool) {
	for _, r := range AccessCodeRules {
		if r.Code == code {
			return r, true
		}
	}
	return Rule{}, false
}
func RequiredRules() []Rule {
	out := []Rule{}
	for _, r := range AccessCodeRules {
		if r.Required {
			out = append(out, r)
		}
	}
	return out
}
func ValidateText(v string, max int) bool {
	if len(v) == 0 || len(v) > max {
		return false
	}
	for _, r := range v {
		if r < ' ' || r == '\u007f' {
			return false
		}
	}
	return true
}
func ValidateRecordFields(r Record) []string {
	errs := []string{}
	if !ValidateText(r.Owner, 80) {
		errs = append(errs, "owner")
	}
	if !ValidateText(r.Title, 160) {
		errs = append(errs, "title")
	}
	if !ValidateText(r.CurrentCode(), 128) {
		errs = append(errs, "code")
	}
	if r.CreatedAt < 0 {
		errs = append(errs, "created")
	}
	if r.UpdatedAt < r.CreatedAt {
		errs = append(errs, "updated")
	}
	return errs
}
