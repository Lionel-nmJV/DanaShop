package merchant

type CustomError struct {
	ErrorCode  int
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(errorCode int, statusCode int, message string) *CustomError {
	return &CustomError{
		ErrorCode:  errorCode,
		StatusCode: statusCode,
		Message:    message,
	}
}
