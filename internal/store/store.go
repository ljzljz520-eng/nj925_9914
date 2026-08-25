package store

import (
	"errors"
	"go.etcd.io/bbolt"
	"path/filepath"
	"privatealbum/internal/model"
	"sort"
	"strings"
)

var buckets = [][]byte{[]byte("records"), []byte("audit"), []byte("workflows"), []byte("attachments")}

type Store struct {
	db   *bbolt.DB
	path string
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func (s *Store) PutRecord(r model.Record) error {
	if err := r.Validate(); err != nil {
		return err
	}
	b, e := model.Encode(r)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[0]).Put([]byte(r.ID), b) })
}
func (s *Store) GetRecord(id string) (model.Record, error) {
	var r model.Record
	err := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(buckets[0]).Get([]byte(id))
		if v == nil {
			return errors.New("record not found")
		}
		return model.Decode(v, &r)
	})
	return r, err
}
func (s *Store) ListRecords() ([]model.Record, error) {
	out := []model.Record{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[0]).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := model.Decode(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, err
}
func (s *Store) DeleteRecord(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[0]).Delete([]byte(id)) })
}
func (s *Store) PutAudit(a model.AuditEvent) error {
	b, e := model.Encode(a)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[1]).Put([]byte(a.ID), b) })
}
func (s *Store) ListAudit(recordID string) ([]model.AuditEvent, error) {
	out := []model.AuditEvent{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[1]).ForEach(func(_, v []byte) error {
			var a model.AuditEvent
			if e := model.Decode(v, &a); e != nil {
				return e
			}
			if a.RecordID == recordID {
				out = append(out, a)
			}
			return nil
		})
	})
	sort.Slice(out, func(i, j int) bool { return out[i].At < out[j].At })
	return out, err
}
func (s *Store) PutWorkflow(w model.Workflow) error {
	b, e := model.Encode(w)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[2]).Put([]byte(w.ID), b) })
}
func (s *Store) PutAttachment(a model.Attachment) error {
	b, e := model.Encode(a)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[3]).Put([]byte(a.ID), b) })
}
func Match(r model.Record, q string) bool {
	q = strings.ToLower(strings.TrimSpace(q))
	return q == "" || strings.Contains(strings.ToLower(r.Owner), q) || strings.Contains(strings.ToLower(r.Title), q)
}
