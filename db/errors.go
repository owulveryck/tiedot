package db

import (
	"errors"
)

var (
	// ErrAlreadyIndexed istriggered when trying to index a path that is already indexed
	ErrAlreadyIndexed = errors.New("Path already indexed")
)
