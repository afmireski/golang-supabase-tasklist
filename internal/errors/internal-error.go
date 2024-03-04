package errors

type InternalError struct {
	Message  string `json:"error_message"`
	HttpCode int    `json:"http_code"`
}

func (e *InternalError) Error() string {
	return e.Message
}

func NewInternalError(message string, httpCode int) *InternalError {
	return &InternalError{message, httpCode}
}
