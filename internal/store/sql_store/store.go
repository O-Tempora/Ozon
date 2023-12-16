package sqlstore

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type SqlStore struct {
	Db *pgx.Conn
}

func CreateSqlStore(port int, host, user, password, base string) (*SqlStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, base)
	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("Database connection failed: %w", err)
	}
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Database ping failed: %w", err)
	}

	if _, err := db.Exec(ctx, `create table urls(
		id serial4 primary key,
		url text not null unique,
		short_url varchar(10) not null unique check (length(short_url) = 10)
	)`); err != nil {
		return nil, fmt.Errorf("Table creation failed: %w", err)
	}
	return &SqlStore{Db: db}, nil
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
