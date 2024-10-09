package response

type Status int

// LOS STATUS DEBEN MANTENER SU ORDEN. LOS NUEVOS DEBEN SER AGREGADOS AL FINAL
const (
	InternalServerError Status = iota
	OK
	Unknown
	SearchFailed
	CreateFailed
	DeleteFailed
	UpdateFailed
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
	UsernamesDoNotMatch
)

// LOS STATUS DEBEN MANTENER SU ORDEN. LOS NUEVOS DEBEN SER AGREGADOS AL FINAL
func (s Status) String() string {
	return [...]string{
		"InternalServerError",
		"OK",
		"Unknown",
		"SearchFailed",
		"CreateFailed",
		"DeleteFailed",
		"UpdateFailed",
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
		"UsernamesDoNotMatch",
	}[s]
}

func (s Status) Index() int {
	return int(s)
}
