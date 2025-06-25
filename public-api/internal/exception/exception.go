package exception

type ApiError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (a *ApiError) Error() string {
	return a.Message
}

func NewApiError(message string, statusCode int) *ApiError {
	return &ApiError{Message: message, StatusCode: statusCode}
}
