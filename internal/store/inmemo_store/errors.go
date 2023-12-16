package inmemostore

import "errors"

var (
	errUrlNotExist = errors.New("URL does not exist in store")
	errInvalidURL  = errors.New("invalid URL value in store")
)
