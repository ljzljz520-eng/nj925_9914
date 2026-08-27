package importer

import (
	"strings"
)

func (r Result) Successful() bool { return r.Rejected == 0 && r.Accepted > 0 }
func (r Result) Message() string {
	if r.Rejected == 0 {
		return "all rows accepted"
	}
	return strings.Join(r.Errors, "; ")
}
