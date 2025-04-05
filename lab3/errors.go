package main


import (
	"errors"
)


var (
	ErrItemExists       = errors.New("item already exists")
	ErrItemNotFound     = errors.New("item not found")
	ErrNotImplemented   = errors.New("operation not implemented")
	ErrPermissionDenied = errors.New("permission denied")
	ErrNotDirectory     = errors.New("not a directory")
	ErrIsDirectory      = errors.New("is a directory")
)