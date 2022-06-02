package common

type ErrorResponse struct {
	StatusCode int
	Message string
}

func NewErrorResponse(statusCode int, message string) *ErrorResponse{
	return &ErrorResponse{
		StatusCode: statusCode,
		Message: message,
	}
}