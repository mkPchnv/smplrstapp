package repository

import "errors"

var (
	ErrDuplicate          = errors.New("record already exists")
	ErrNotExist           = errors.New("row does not exist")
	ErrUpdateFailed       = errors.New("update failed")
	ErrForeignKeyNotExist = errors.New("foreign key not exist")
)

const dublicateErrorCode = "23505"
