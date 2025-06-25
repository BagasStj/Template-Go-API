package errors

import "net/http"

// ErrorCode represents error codes
type ErrorCode int

const (
	ErrCodeBadRequest ErrorCode = iota + 1
	ErrCodeUnauthorized
	ErrCodeForbidden
	ErrCodeNotFound
	ErrCodeConflict
	ErrCodeInternalServerError
	ErrCodeValidation
)

// APIError represents an API error
type APIError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Details string    `json:"details,omitempty"`
}

// Error returns the error message
func (e APIError) Error() string {
	return e.Message
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success bool     `json:"success"`
	Error   APIError `json:"error"`
}

// SetResponse returns an error response based on error code
func SetResponse(code ErrorCode) ErrorResponse {
	var message string
	switch code {
	case ErrCodeBadRequest:
		message = "Bad request"
	case ErrCodeUnauthorized:
		message = "Unauthorized"
	case ErrCodeForbidden:
		message = "Forbidden"
	case ErrCodeNotFound:
		message = "Not found"
	case ErrCodeConflict:
		message = "Conflict"
	case ErrCodeInternalServerError:
		message = "Internal server error"
	case ErrCodeValidation:
		message = "Validation error"
	default:
		message = "Unknown error"
	}

	return ErrorResponse{
		Success: false,
		Error: APIError{
			Code:    code,
			Message: message,
		},
	}
}

// SetResponseWithDetails returns an error response with custom details
func SetResponseWithDetails(code ErrorCode, details string) ErrorResponse {
	response := SetResponse(code)
	response.Error.Details = details
	return response
}

// GetHTTPStatusCode returns the HTTP status code for the error code
func GetHTTPStatusCode(code ErrorCode) int {
	switch code {
	case ErrCodeBadRequest:
		return http.StatusBadRequest
	case ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrCodeForbidden:
		return http.StatusForbidden
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeValidation:
		return http.StatusUnprocessableEntity
	case ErrCodeInternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
