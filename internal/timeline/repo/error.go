package repo

import "errors"

var (
	ErrNotFound      = errors.New("record not found")
	ErrAlreadyExists = errors.New("record already exists")
	ErrAborted       = errors.New("operation aborted")
)
