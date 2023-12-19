package response

type Status int

// LOS STATUS DEBEN MANTENER SU ORDEN. LOS NUEVOS DEBEN SER AGREGADOS AL FINAL
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

// LOS STATUS DEBEN MANTENER SU ORDEN. LOS NUEVOS DEBEN SER AGREGADOS AL FINAL
func (s Status) String() string {
	return [...]string{
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
	}[s]
}

func (s Status) Index() int {
	return int(s)
}
