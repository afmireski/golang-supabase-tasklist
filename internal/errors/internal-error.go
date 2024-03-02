package errors

type InternalError struct {
	message string
	httpCode int
}

func (e *InternalError) Error() string {
	return e.message
}

func (e *InternalError) HttpCode() int {
	return e.httpCode
}

func NewInternalError(message string, httpCode int) *InternalError {
	return &InternalError{message, httpCode}
}