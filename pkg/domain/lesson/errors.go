package lesson

import "errors"

var (
	ErrRepo = errors.New("unable to handle repository request")
)
