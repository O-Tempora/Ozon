package inmemostore

import (
	"context"
	"sync"
)

type InmemStore struct {
	Map *sync.Map
}

func (st *InmemStore) SaveShortenedURL(ctx context.Context, url, shortURL string) error {
	st.Map.Store(shortURL, url)
	return nil
}

func (st *InmemStore) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	url, ok := st.Map.Load(shortURL)
	if !ok {
		return "", errUrlNotExist
	}
	urlVal, ok := url.(string)
	if !ok {
		return "", errInvalidURL
	}
	return urlVal, nil
}
