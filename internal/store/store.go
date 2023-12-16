package store

import (
	"context"
)

type Store interface {
	SaveShortenedURL(ctx context.Context, url, shortURL string) error
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
}
