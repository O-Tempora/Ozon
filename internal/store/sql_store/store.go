package sqlstore

import "github.com/jackc/pgx/v5"

type SqlStore struct {
	Db *pgx.Conn
}

func (st *SqlStore) SaveShortenedURL(url string) error {
	return nil
}

func (st *SqlStore) GetOriginalURL(shortURL string) (string, error) {
	return "", nil
}
