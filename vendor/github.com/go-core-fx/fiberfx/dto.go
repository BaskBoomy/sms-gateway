package fiberfx

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
	Details any    `json:"details,omitempty"`
}

func NewErrorResponse(message string, code int, details any) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Code:    code,
		Details: details,
	}
}
