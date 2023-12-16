package store

import (
	"context"
	"fmt"
	"sync"
	"time"

	inmemostore "github.com/O-Tempora/Ozon/internal/store/inmemo_store"
	sqlstore "github.com/O-Tempora/Ozon/internal/store/sql_store"
	"github.com/jackc/pgx/v5"
)

type Store interface {
	ShortenURL(url string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
}

func CreateSqlStore(port int, host, user, password, base string) (*sqlstore.SqlStore, error) {
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
		url text not null,
		short_url varchar(10) not null check (length(short_url) = 10)
	)`); err != nil {
		return nil, fmt.Errorf("Table creation failed: %w", err)
	}
	return &sqlstore.SqlStore{Db: db}, nil
}

func CreateInmemoStore() *inmemostore.InmemStore {
	return &inmemostore.InmemStore{
		Store: &sync.Map{},
	}
}
