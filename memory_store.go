package mls

import (
	"sync"
)

type MemoryStore struct {
	mu      sync.RWMutex
	entries map[string][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		entries: make(map[string][]byte, 0),
	}
}

func (s *MemoryStore) Upsert(key, value []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries[string(key)] = value
	return nil
}

func (s *MemoryStore) Get(key []byte) ([]byte, error) {
	if v, ok := s.entries[string(key)]; ok {
		return v, nil
	}
	return nil, nil
}

func (s *MemoryStore) Remove(key []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := string(key)
	if _, ok := s.entries[k]; ok {
		delete(s.entries, k)
		return nil
	}
	return nil
}

func (s *MemoryStore) Clear() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries = nil
	s.entries = make(map[string][]byte, 0)
	return nil
}
