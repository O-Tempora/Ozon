package inmemostore

import "errors"

var (
	errUrlNotExist           = errors.New("Short URL does not exist in store")
	errInvalidURL            = errors.New("invalid URL value in store")
	errUrlAlreadyExists      = errors.New("URL already exists")
	errShortUrlAlreadyExists = errors.New("Short URL already exists")
)
