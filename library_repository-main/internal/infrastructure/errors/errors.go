package errors

const (
	ResponderEncodeError        = "json output: error during data serialization"
	ResponderStatusUnauthorized = "http response: unauthorized"
	ResponderErrorUnauthorized  = "unauthorized: error during data serialization"
	ResponderStatusBadRequest   = "http response: bad request"
	ResponderErrorBadRequest    = "bad request: error during data serialization"
	ResponderStatusForbidden    = "http response: forbidden"
	ResponderErrorForbidden     = "forbidden: error during data serialization"
	ResponderStatusInternal     = "http response: internal"
	ResponderErrorInternal      = "internal: error during data serialization"
)

// Common
const (
	HandlerParamsError = "error during url params parsing"
	HandlerFileError   = "file retrieval error"
)

const (
	ServiceLoginNotFound  = "error: this user is not registered"
	ServiceLoginAlreadyIn = "error: you're already logged in"
	ServiceLogout         = "error: you didn't come in)"
	ServiceGetByIDError   = "error during data selection by ID"
	ServiceIDEncodeError  = "error during id parsing"
)

const (
	StorageAlreadyExistsError  = "error: entity already created"
	StorageAlreadyDeletedError = "error: entity already deleted"
)
