package merchant

type customError struct {
	ErrorCode  int
	StatusCode int
	Message    string
}

func (e *customError) Error() string {
	return e.Message
}

func newCustomError(errorCode int, statusCode int, message string) *customError {
	return &customError{
		ErrorCode:  errorCode,
		StatusCode: statusCode,
		Message:    message,
	}
}
