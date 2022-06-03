package common

//ErrorResponse - type for response error with status code
type ErrorResponse struct {
	StatusCode int
	Message string
}

//NewErrorResponse - constructor for object of ErrorResponse 
func NewErrorResponse(statusCode int, message string) *ErrorResponse{
	return &ErrorResponse{
		StatusCode: statusCode,
		Message: message,
	}
}