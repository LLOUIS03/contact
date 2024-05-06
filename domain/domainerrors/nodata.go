package domainerrors

// NoDataError is an error type for no data errors
type NoDataError struct {
	base
}

// Message returns the error message
func (e NoDataError) Message() string {
	return e.message
}

// NewNoDataError creates a new NoDataError
func NewNoDataError(innerError error) NoDataError {
	return NoDataError{
		base: base{
			message:    "no data found",
			innerError: innerError,
		},
	}
}
