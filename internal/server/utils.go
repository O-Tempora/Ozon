package server

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net/url"
	"strings"
)

const (
	base          uint64 = 63
	shortenedSize        = 10
)

var (
	alphabet = [base]rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '_',
	}
)

func encodeURL(num uint64) (string, error) {
	var b strings.Builder
	b.Grow(shortenedSize)
	for ; num > 0 && b.Len() < shortenedSize; num /= base {
		_, err := b.WriteRune(alphabet[num%base])
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

func randomizeUint64() (uint64, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b[:]), nil
}

func shortenURL(longURL string) (string, error) {
	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		return "", fmt.Errorf("Failed to validate URL: %w", err)
	}
	u, err := randomizeUint64()
	if err != nil {
		return "", err
	}
	shortURL, err := encodeURL(u)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}
