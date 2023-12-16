package inmemostore

import "sync"

type InmemStore struct {
	Store *sync.Map
}

func (st *InmemStore) SaveShortenedURL(url string) error {
	return nil
}

func (st *InmemStore) GetOriginalURL(shortURL string) (string, error) {
	return "", nil
}
