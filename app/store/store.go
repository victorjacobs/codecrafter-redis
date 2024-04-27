package store

type Store struct {
	store map[string]string
}

func NewStore() *Store {
	return &Store{
		store: map[string]string{},
	}
}

func (s *Store) Set(key, value string) error {
	s.store[key] = value

	return nil
}

func (s *Store) Get(key string) (string, bool) {
	if value, present := s.store[key]; present {
		return value, true
	} else {
		return "", false
	}
}
