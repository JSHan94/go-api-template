package lib

import "errors"

var (
	ErrInvalidRequest = errors.New("worng uri for the request")
	ErrNotFound       = errors.New("no hits for the request")
)
