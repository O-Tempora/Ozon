package store

type Store interface {
	ShortenURL(url string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
}
