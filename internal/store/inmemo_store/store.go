package inmemostore

import (
	"context"
	"sync"
)

type InmemStore struct {
	pairs *sync.Map
	urls  *sync.Map
}

// pairs - хранит ключ-shortURL : значение-url
// urls - хранит ключ-уже записанные URL : значение-struct{}{}
// Сделано для производительности, чтобы каждый раз
// не итерировать по мапе pairs и не искать, существует ли
// запись с таким значением
func CreateInmemoStore() *InmemStore {
	return &InmemStore{
		pairs: &sync.Map{},
		urls:  &sync.Map{},
	}
}

func (st *InmemStore) SaveShortenedURL(ctx context.Context, url, shortURL string) error {
	if _, ok := st.urls.Load(url); ok {
		return errUrlAlreadyExists
	}
	if _, ok := st.pairs.Load(shortURL); ok {
		return errShortUrlAlreadyExists
	}
	st.urls.Store(url, struct{}{})
	st.pairs.Store(shortURL, url)
	return nil
}

func (st *InmemStore) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	url, ok := st.pairs.Load(shortURL)
	if !ok {
		return "", errUrlNotExist
	}
	urlVal, ok := url.(string)
	if !ok {
		return "", errInvalidURL
	}
	return urlVal, nil
}
