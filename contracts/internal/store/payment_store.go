package store

import "errors"

// ErrNotFound indicates that a requested payment does not exist.
var ErrNotFound = errors.New("payment not found")
