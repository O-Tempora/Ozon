package server

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/O-Tempora/Ozon/config"
	"github.com/O-Tempora/Ozon/internal/api/shortener_v1"
	"github.com/O-Tempora/Ozon/internal/store"
	"github.com/rs/zerolog"
)

type Server struct {
	Logger zerolog.Logger
	Store  store.Store
	shortener_v1.UnimplementedShortenerServiceServer
}

func initLogger(src io.Writer) zerolog.Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        src,
		NoColor:    false,
		TimeFormat: time.ANSIC,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatTimestamp: func(i interface{}) string {
			t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", i))
			return t.Format(time.RFC1123)
		},
	}).With().Timestamp().Logger().Level(zerolog.InfoLevel)
	return logger
}

func initStore(useDb bool, cf *config.Config) (store.Store, error) {
	if useDb {
		store, err := store.CreateSqlStore(cf.DbPort, cf.DbHost, cf.DbUser, cf.DbPass, cf.DbName)
		if err != nil {
			return nil, fmt.Errorf("Failed to initialize database: %w", err)
		}
		return store, nil
	}
	store := store.CreateInmemoStore()
	return store, nil
}

func CreateServer(useDb bool, cf *config.Config) (*Server, error) {
	s := &Server{
		Logger: initLogger(os.Stdout),
	}
	store, err := initStore(useDb, cf)
	if err != nil {
		return nil, fmt.Errorf("Failed to create server: %w", err)
	}
	s.Store = store

	return s, nil
}
