package inmemostore

import "sync"

type InmemStore struct {
	Store *sync.Map
}

func (st *InmemStore) ShortenURL(url string) (string, error) {
	return "", nil
}

func (st *InmemStore) GetOriginalURL(shortURL string) (string, error) {
	return "", nil
}
