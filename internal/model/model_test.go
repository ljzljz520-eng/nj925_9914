package model

import "testing"

func TestRecordValidation(t *testing.T) {
	r := Record{ID: "1", Owner: "a", Title: "t", AccessCodes: []string{"x"}, Status: Pending}
	if e := r.Validate(); e != nil {
		t.Fatal(e)
	}
	r.Status = "bad"
	if r.Validate() == nil {
		t.Fatal("expected status error")
	}
}
