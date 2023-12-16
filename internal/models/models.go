package models

import "time"

type Pair struct {
	LongURL  string
	ShortURL string
}

type PairWithExparation struct {
	LongURL        string
	ShortURL       string
	ExpirationTime time.Duration
}
