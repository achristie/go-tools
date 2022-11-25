package store

import (
	"encoding/gob"
	"io"
	"os"
)

type Store struct {
	stream io.ReadWriteCloser
	path   string
}

func Open(pathname string) *Store {
	return &Store{path: pathname}
}

func (s *Store) Close() error {
	if s.stream == nil {
		return nil
	}
	return s.stream.Close()
}

func (s *Store) Save(v any) error {
	if s.stream == nil {
		f, err := os.Create(s.path)
		if err != nil {
			return err
		}
		s.stream = f
	}
	return gob.NewEncoder(s.stream).Encode(v)
}

func (s *Store) Load(v any) error {
	if s.stream == nil {
		f, err := os.Open(s.path)
		if err != nil {
			return err
		}
		s.stream = f
	}
	return gob.NewDecoder(s.stream).Decode(v)
}
