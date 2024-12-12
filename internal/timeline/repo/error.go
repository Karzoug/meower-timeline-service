package repo

import "errors"

var (
	ErrKeyNotFound   = errors.New("record key not found")
	ErrValueNotFound = errors.New("record value not found")
	ErrAlreadyExists = errors.New("record already exists")
	ErrAborted       = errors.New("operation aborted")
)
