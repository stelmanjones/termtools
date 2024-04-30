package errors

import "errors"

var (
	ErrKeyNotFound   = errors.New("key not found")
	ErrKeyExists     = errors.New("key exists")
	ErrInvalidKey    = errors.New("invalid key")
	ErrInvalidJSON  = errors.New("invalid json")
	ErrMissingValue = errors.New("missing value")
	ErrInvalidValue  = errors.New("invalid value")
	ErrTableNotFound = errors.New("table not found")
	ErrTableFull    = errors.New("table is at max capacity")
)
