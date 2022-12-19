package domain

import "errors"

var (

	// ErrInvalidPort means that the given port data is invalid
	ErrInvalidPort = errors.New("invalid port data")
	// ErrNoRecords means that no record was found on the database.
	ErrNoRecords = errors.New("no records were found")
	//ErrRecordAlreadyExists means that the given record already exists in the database.
	ErrRecordAlreadyExists = errors.New("informed record already exists")
)
