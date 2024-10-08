package response

type Status int

// LOS STATUS DEBEN MANTENER SU ORDEN. LOS NUEVOS DEBEN SER AGREGADOS AL FINAL
const (
	InternalServerError Status = iota
	OK
	Unknown
	BadRequest
	Unauthorized
	Forbidden
	MethodNotAllowed
	Conflict
	ServiceUnavailable
	Timeout
	UnsupportedMediaType
	TooManyRequests
	NotAcceptable
	PreconditionFailed
	Gone
	UnprocessableEntity
	RequestTimeout
	LengthRequired
	ExpectationFailed
	PayloadTooLarge
	UriTooLong
	NotImplemented
	BadGateway
	ServiceTimeout
	InsufficientStorage
	LoopDetected
	NetworkAuthenticationRequired
	UnavailableForLegalReasons
	ProxyAuthenticationRequired
	FailedDependency
	NotFound
	SearchFailed
	CreateFailed
	UpdateFailed
	DeleteFailed
	InvalidRequest
	ResourceAlreadyExists
	ResourceLocked
	FetchFailed
	InvalidParameter
	DatabaseError
	ValidationFailed
	PermissionDenied
	UnexpectedError
	TimeoutError
	MethodNotSupported
	PatchFailed
	FetchNotModified
	UpdateConflict
	DeleteConflict
	ExpiredSession
	InvalidCredentials
	ResourceNotModified
	LockedResource
	UnauthorizedAccess
	InsufficientPermissions
	UrlNotFound
	ForbiddenAction
	IncompleteData
	SessionTimeout
	ApiVersionNotSupported
	InactiveAccount
	BlockedAccount
	DBInitError
	DBQueryError
	DBExecutionError
	DBScanError
	DBRowsError
	DBRowsAffectedError
	NoContent
	FormatInvalid
	DecodeError
	RequestEntityTooLarge
	ServiceTemporarilyUnavailable
	PreconditionRequired
	TooManyConnections
	InvalidApiKey
	InactiveSession
	ResourceNotAvailable
	UnsupportedVersion
	FeatureNotAvailable
	RequestRateLimited
	OperationNotAllowed
)

// LOS STATUS DEBEN MANTENER SU ORDEN. LOS NUEVOS DEBEN SER AGREGADOS AL FINAL
func (s Status) String() string {
	return [...]string{
		 "InternalServerError",
		 "OK",
		 "Unknown",
		 "BadRequest",
		 "Unauthorized",
		 "Forbidden",
		 "MethodNotAllowed",
		 "Conflict",
		 "ServiceUnavailable",
		 "Timeout",
		 "UnsupportedMediaType",
		 "TooManyRequests",
		 "NotAcceptable",
		 "PreconditionFailed",
		 "Gone",
		 "UnprocessableEntity",
		 "RequestTimeout",
		 "LengthRequired",
		 "ExpectationFailed",
		 "PayloadTooLarge",
		 "UriTooLong",
		 "NotImplemented",
		 "BadGateway",
		 "ServiceTimeout",
		 "InsufficientStorage",
		 "LoopDetected",
		 "NetworkAuthenticationRequired",
		 "UnavailableForLegalReasons",
		 "ProxyAuthenticationRequired",
		 "FailedDependency",
		 "NotFound",
		 "SearchFailed",
		 "CreateFailed",
		 "UpdateFailed",
		 "DeleteFailed",
		 "InvalidRequest",
		 "ResourceAlreadyExists",
		 "ResourceLocked",
		 "FetchFailed",
		 "InvalidParameter",
		 "DatabaseError",
		 "ValidationFailed",
		 "PermissionDenied",
		 "UnexpectedError",
		 "TimeoutError",
		 "MethodNotSupported",
		 "PatchFailed",
		 "FetchNotModified",
		 "UpdateConflict",
		 "DeleteConflict",
		 "ExpiredSession",
		 "InvalidCredentials",
		 "ResourceNotModified",
		 "LockedResource",
		 "UnauthorizedAccess",
		 "InsufficientPermissions",
		 "UrlNotFound",
		 "ForbiddenAction",
		 "IncompleteData",
		 "SessionTimeout",
		 "ApiVersionNotSupported",
		 "InactiveAccount",
		 "BlockedAccount",
		 "DBInitError",
		 "DBQueryError",
		 "DBExecutionError",
		 "DBScanError",
		 "DBRowsError",
		 "DBRowsAffectedError",
		 "NoContent",
		 "FormatInvalid",
		 "DecodeError",
		 "RequestEntityTooLarge",
		 "ServiceTemporarilyUnavailable",
		 "PreconditionRequired",
		 "TooManyConnections",
		 "InvalidApiKey",
		 "InactiveSession",
		 "ResourceNotAvailable",
		 "UnsupportedVersion",
		 "FeatureNotAvailable",
		 "RequestRateLimited",
		 "OperationNotAllowed",
	}[s]
}

func (s Status) Index() int {
	return int(s)
}
