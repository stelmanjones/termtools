package errors

import "errors"





var (
	ErrKeyNotFound = errors.New("key not found")
	ErrKeyExists   = errors.New("key exists")
	ErrInvalidKey  = errors.New("invalid key")
	ErrInvalidValue = errors.New("invalid value")
	ErrTableNotFound = errors.New("table not found")
)