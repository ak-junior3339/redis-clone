package main

import (
	"sync"
	"time"
)

type Store struct {
	mu     sync.RWMutex
	data   map[string]string
	expiry map[string]time.Time
}

func NewStore() *Store {
	return &Store{
		data:   make(map[string]string),
		expiry: make(map[string]time.Time),
	}
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// check if key has expired
	if expireAt, ok := s.expiry[key]; ok {
		if time.Now().After(expireAt) {
			return "", false // expired, act like it doesn't exist
		}
	}
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
func (s *Store) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.data[key]
	return ok
}

func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	keys := make([]string, 0)
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *Store) Flush() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[string]string)
}
func (s *Store) Expire(key string, seconds int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// time.Now().Add() sets the expiry time in the future
	s.expiry[key] = time.Now().Add(time.Duration(seconds) * time.Second)
}

func (s *Store) TTL(key string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	expireAt, ok := s.expiry[key]
	if !ok {
		return -1 // no expiry set
	}
	remaining := time.Until(expireAt).Seconds()
	if remaining < 0 {
		return -1 // already expired
	}
	return int(remaining)
}
