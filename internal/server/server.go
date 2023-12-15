package server

import (
	"github.com/O-Tempora/Ozon/internal/api/shortener_v1"
	"github.com/O-Tempora/Ozon/internal/store"
	"github.com/rs/zerolog"
)

type Server struct {
	Logger zerolog.Logger
	Store  store.Store
	shortener_v1.UnimplementedShortenerServiceServer
}
