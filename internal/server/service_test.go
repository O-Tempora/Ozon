package server

import (
	"context"
	"os"
	"testing"

	"github.com/O-Tempora/Ozon/internal/api/shortener_v1"
	inmemostore "github.com/O-Tempora/Ozon/internal/store/inmemo_store"
)

var server *Server

func TestMain(m *testing.M) {
	server = &Server{
		Logger: initLogger(os.Stdout),
		Store:  inmemostore.CreateInmemoStore(),
	}
	m.Run()
}

func TestEncodeEmptyStr(t *testing.T) {
	var err error
	tc := struct {
		number uint64
		want   string
		got    string
	}{0, "", ""}

	tc.got, err = encodeURL(tc.number)
	if err != nil {
		t.Fatalf("encoding error: %s", err.Error())
	}
	if tc.want != tc.got {
		t.Errorf("Invalid encode result for %d: want %s, got %s", tc.number, tc.want, tc.got)
	}
}

func TestEncodeNormal(t *testing.T) {
	var err error
	tt := []struct {
		number uint64
		want   string
		got    string
	}{
		{12, "M", ""},
		{69, "GB", ""},
		{420, "qG", ""},
	}
	for _, tc := range tt {
		tc.got, err = encodeURL(tc.number)
		if err != nil {
			t.Error(err)
			continue
		}
		if tc.want != tc.got {
			t.Errorf("Invalid encode result for %d: want %s, got %s", tc.number, tc.want, tc.got)
		}
	}
}

func TestShortenUrlInvalidURL(t *testing.T) {
	tc := struct {
		url    string
		gotErr error
	}{"AbsolutelyNotValidURL123", nil}

	_, tc.gotErr = shortenURL(tc.url)
	if tc.gotErr == nil {
		t.Errorf("Invalid shortener error result for %s: want not nil", tc.url)
	}
}

func TestGetShortenedUrl(t *testing.T) {
	tt := []struct {
		url   string
		short *shortener_v1.ShortenedURL
		err   error
	}{
		{"https://www.scala-lang.org/#processing-data", nil, nil},
		{"https://elixirschool.com/en", nil, nil},
		{"https://yandex.ru/yaintern/int_01", nil, nil},
	}
	for _, tc := range tt {
		tc.short, tc.err = server.GetShortenedURL(context.Background(), &shortener_v1.LongURL{LongURL: tc.url})
		if tc.err != nil {
			t.Error(tc.err)
			continue
		}
		if len(tc.short.ShortURL) != 10 {
			t.Errorf("Invalid short url for %s: want length %d, got %d", tc.url, 10, len(tc.short.ShortURL))
		}
	}
	t.Cleanup(func() {
		server.Store = inmemostore.CreateInmemoStore()
	})
}
