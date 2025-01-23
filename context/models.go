package main

type Store interface {
	Fetch() string
	Cancel()
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func (s *SpyStore) Cancel() {
    s.cancelled = true
}

