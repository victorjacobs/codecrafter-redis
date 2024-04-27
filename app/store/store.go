package store

import "time"

// Store is a key-value store. It is not thread-safe.
type Store struct {
	store map[string]*value
}

type value struct {
	value     string
	expiresAt int64
}

func NewStore() *Store {
	return &Store{
		store: map[string]*value{},
	}
}

func (s *Store) Set(k, v string) error {
	return s.SetWithExpiry(k, v, 0)
}

func (s *Store) Delete(k string) error {
	delete(s.store, k)

	return nil
}

func (s *Store) SetWithExpiry(k, v string, expiresAt int64) error {
	s.store[k] = &value{
		value:     v,
		expiresAt: expiresAt,
	}

	return nil
}

func (s *Store) Get(key string) (string, bool) {
	if val, present := s.store[key]; present {
		if val.expiresAt != 0 && val.expiresAt < time.Now().UnixNano() {
			s.Delete(key)
			return "", false
		}

		return val.value, true
	} else {
		return "", false
	}
}
