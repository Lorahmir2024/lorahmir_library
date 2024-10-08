package response

import (
	"errors"
)

type Status int

const (
	InternalServerError Status = iota
	Unknown
	DBInitError
	DBQueryError
	DBExecutionError
	DBScanError
	DBRowsError
	DBRowsAffectedError
	NoContent
	BadRequest
	Conflict
	FormatInvalid
	DecodeError
	Unauthorized
	RequestEntityTooLarge
	UnsupportedMediaType
	SearchSuccess
	SearchFailed
	CreateSuccess
	CreateFailed
	DeleteSuccess
	DeleteFailed
	UpdateSuccess
	UpdateFailed
)

var statusNames = []string{
	"InternalServerError",
	"Unknown",
	"DBInitError",
	"DBQueryError",
	"DBExecutionError",
	"DBScanError",
	"DBRowsError",
	"DBRowsAffectedError",
	"NoContent",
	"BadRequest",
	"Conflict",
	"FormatInvalid",
	"DecodeError",
	"Unauthorized",
	"RequestEntityTooLarge",
	"UnsupportedMediaType",
	"SearchSuccess",
	"SearchFailed",
	"CreateSuccess",
	"CreateFailed",
	"DeleteSuccess",
	"DeleteFailed",
	"UpdateSuccess",
	"UpdateFailed",
}

func (s Status) String() string {
	if s < InternalServerError || s > UpdateFailed {
		return "Invalid Status"
	}
	return statusNames[s]
}

func (s Status) Index() int {
	return int(s)
}

func (s Status) Valid() error {
	if s < InternalServerError || s > UpdateFailed {
		return errors.New("invalid status")
	}
	return nil
}
