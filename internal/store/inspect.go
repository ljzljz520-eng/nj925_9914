package store

import (
	"fmt"
	"go.etcd.io/bbolt"
	"privatealbum/internal/model"
)

type Snapshot struct {
	Records     int
	Audits      int
	Workflows   int
	Attachments int
}

func (s *Store) Snapshot() (Snapshot, error) {
	out := Snapshot{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		for _, pair := range []struct {
			dst *int
			b   []byte
		}{{&out.Records, buckets[0]}, {&out.Audits, buckets[1]}, {&out.Workflows, buckets[2]}, {&out.Attachments, buckets[3]}} {
			n := 0
			e := tx.Bucket(pair.b).ForEach(func(_, v []byte) error {
				if v != nil {
					n++
				}
				return nil
			})
			if e != nil {
				return e
			}
			*pair.dst = n
		}
		return nil
	})
	return out, err
}
func (s *Store) UpsertRecords(rs []model.Record) error {
	for _, r := range rs {
		if e := s.PutRecord(r); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) ReplaceRecord(id string, fn func(model.Record) (model.Record, error)) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	next, e := fn(r)
	if e != nil {
		return e
	}
	if next.ID != id {
		return fmt.Errorf("identifier changed")
	}
	return s.PutRecord(next)
}
func (s *Store) Transaction(fn func(*bbolt.Tx) error) error { return s.db.Update(fn) }
func (s *Store) Exists(id string) bool                      { _, e := s.GetRecord(id); return e == nil }
