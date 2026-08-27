package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Record struct {
	ID, Owner, Title     string
	AccessCodes          []string
	Status               string
	CreatedAt, UpdatedAt int64
}
type AuditEvent struct {
	ID, RecordID, Actor, Action, Detail string
	At                                  int64
}
type Workflow struct{ ID, RecordID, Kind, State, Payload string }
type Attachment struct {
	ID, RecordID, Name, Checksum string
	Size                         int64
}

const (
	Pending  = "pending"
	Approved = "approved"
	Rejected = "rejected"
	Archived = "archived"
)

func (r Record) Validate() error {
	if strings.TrimSpace(r.ID) == "" {
		return errors.New("id required")
	}
	if strings.TrimSpace(r.Owner) == "" {
		return errors.New("owner required")
	}
	if strings.TrimSpace(r.Title) == "" {
		return errors.New("title required")
	}
	if len(r.AccessCodes) == 0 || strings.TrimSpace(r.AccessCodes[0]) == "" {
		return errors.New("access code required")
	}
	switch r.Status {
	case Pending, Approved, Rejected, Archived:
	default:
		return fmt.Errorf("invalid status %q", r.Status)
	}
	return nil
}
func (r Record) CurrentCode() string {
	if len(r.AccessCodes) == 0 {
		return ""
	}
	return r.AccessCodes[len(r.AccessCodes)-1]
}
func (r Record) PublicSummary() map[string]string {
	return map[string]string{"id": r.ID, "owner": r.Owner, "title": r.Title, "status": r.Status, "code": r.CurrentCode()}
}
func NewID(seed string) string     { h := sha256.Sum256([]byte(seed)); return hex.EncodeToString(h[:8]) }
func Encode(v any) ([]byte, error) { return json.Marshal(v) }
func Decode(b []byte, v any) error { return json.Unmarshal(b, v) }
