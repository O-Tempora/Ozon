package server

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/O-Tempora/Ozon/internal/api/shortener_v1"
)

func (s *Server) GetShortenedURL(ctx context.Context, longURL *shortener_v1.LongURL) (*shortener_v1.ShortenedURL, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(500*time.Millisecond))
	defer cancel()

	_, err := url.ParseRequestURI(longURL.LongURL)
	if err != nil {
		return nil, fmt.Errorf("Invalid URL: %w", err)
	}

	shortURL, err := shortenURL()
	if err != nil {
		return nil, fmt.Errorf("Failed to shorten URL: %w", err)
	}
	if err = s.Store.SaveShortenedURL(shortURL); err != nil {
		return nil, fmt.Errorf("Failed to save URL: %w", err)
	}
	return &shortener_v1.ShortenedURL{ShortURL: shortURL}, nil
}

func (s *Server) GetURL(ctx context.Context, shortURL *shortener_v1.ShortenedURL) (*shortener_v1.LongURL, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(500*time.Millisecond))
	defer cancel()

	longURL, err := s.Store.GetOriginalURL(shortURL.ShortURL)
	if err != nil {
		err = fmt.Errorf("Failed to get original URL: %w", err)
		s.Logger.Error().Msg(err.Error())
		return nil, err
	}
	s.Logger.Info().Msgf("Short: %s, original: %s", shortURL.ShortURL, longURL)
	return &shortener_v1.LongURL{LongURL: longURL}, nil
}
