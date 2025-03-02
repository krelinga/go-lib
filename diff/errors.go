package diff

import "errors"

var (
	ErrFieldNotFound = errors.New("field not found")
	ErrFieldNotExported = errors.New("field not exported")
	ErrFieldWrongType = errors.New("field has wrong type")
	ErrMethodNotFound = errors.New("method not found")
	ErrMethodNotExported = errors.New("method not exported")
	ErrMethodWrongType = errors.New("method has wrong type")
)