package server

import (
	"context"
	"fmt"
	"time"

	"github.com/O-Tempora/Ozon/internal/api/shortener_v1"
)

func (s *Server) GetShortenedURL(ctx context.Context, longURL *shortener_v1.LongURL) (*shortener_v1.ShortenedURL, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(1*time.Second))
	defer cancel()
	shortURL, err := shortenURL(longURL.LongURL)
	if err != nil {
		err = fmt.Errorf("Failed to shorten URL: %w", err)
		s.Logger.Error().Msg(err.Error())
		return nil, err
	}
	if err = s.Store.SaveShortenedURL(ctx, longURL.LongURL, shortURL); err != nil {
		err = fmt.Errorf("Failed to save URL: %w", err)
		s.Logger.Error().Msg(err.Error())
		return nil, err
	}
	s.Logger.Info().Msgf("Successfuly created new short URL: %s", shortURL)
	return &shortener_v1.ShortenedURL{ShortURL: shortURL}, nil
}

func (s *Server) GetURL(ctx context.Context, shortURL *shortener_v1.ShortenedURL) (*shortener_v1.LongURL, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(1*time.Second))
	defer cancel()

	longURL, err := s.Store.GetOriginalURL(ctx, shortURL.ShortURL)
	if err != nil {
		err = fmt.Errorf("Failed to get original URL: %w", err)
		s.Logger.Error().Msg(err.Error())
		return nil, err
	}
	s.Logger.Info().Msgf("Successfuly got URL by %s", shortURL.ShortURL)
	return &shortener_v1.LongURL{LongURL: longURL}, nil
}
