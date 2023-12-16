package sqlstore

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type SqlStore struct {
	Db *pgx.Conn
}

func (st *SqlStore) SaveShortenedURL(ctx context.Context, url, shortURL string) error {
	_, err := st.Db.Exec(
		ctx,
		`insert into urls (url, short_url) values ($1, $2)`,
		url, shortURL,
	)
	if err != nil {
		return fmt.Errorf("Failed to save url %s: %w", url, err)
	}
	return nil
}

func (st *SqlStore) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	var url string
	if err := st.Db.QueryRow(
		ctx,
		`select url from urls where short_url = $1`,
		shortURL,
	).Scan(&url); err != nil {
		return "", fmt.Errorf("Failed to get URL by %s: %w", shortURL, err)
	}
	return url, nil
}
