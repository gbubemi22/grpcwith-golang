package errors

// BaseError is a struct that implements the error interface
type BaseError struct {
	Message        string
	HTTPStatusCode int
	ErrorCode      int
}

func (e *BaseError) Error() string {
	return e.Message
}

// UnauthorizedError represents a 401 Unauthorized error
type UnauthorizedError struct {
	BaseError
}

func NewUnauthorizedError(message string, errorCode int) *UnauthorizedError {
	return &UnauthorizedError{
		BaseError: BaseError{
			Message:        message,
			HTTPStatusCode: 401,
			ErrorCode:      errorCode,
		},
	}
}

// BadRequestError represents a 400 Bad Request error
type BadRequestError struct {
	BaseError
}

func NewBadRequestError(message string, errorCode int) *BadRequestError {
	return &BadRequestError{
		BaseError: BaseError{
			Message:        message,
			HTTPStatusCode: 400,
			ErrorCode:      errorCode,
		},
	}
}

// ConflictError represents a 409 Conflict error
type ConflictError struct {
	BaseError
}

func NewConflictError(message string, errorCode int) *ConflictError {
	return &ConflictError{
		BaseError: BaseError{
			Message:        message,
			HTTPStatusCode: 409,
			ErrorCode:      errorCode,
		},
	}
}

// InternalServerError represents a 500 Internal Server Error
type InternalServerError struct {
	BaseError
}

func NewInternalServerError(message string, errorCode int) *InternalServerError {
	return &InternalServerError{
		BaseError: BaseError{
			Message:        message,
			HTTPStatusCode: 500,
			ErrorCode:      errorCode,
		},
	}
}

// UnauthenticatedError represents a 401 Unauthenticated error
type UnauthenticatedError struct {
	BaseError
}

func NewUnauthenticatedError(message string, errorCode int) *UnauthenticatedError {
	return &UnauthenticatedError{
		BaseError: BaseError{
			Message:        message,
			HTTPStatusCode: 401,
			ErrorCode:      errorCode,
		},
	}
}

// NotFoundError represents a 404 Not Found error
type NotFoundError struct {
	BaseError
}

func NewNotFoundError(message string, errorCode int) *NotFoundError {
	return &NotFoundError{
		BaseError: BaseError{
			Message:        message,
			HTTPStatusCode: 404,
			ErrorCode:      errorCode,
		},
	}
}

// ValidationError represents a 400 Validation error
type ValidationError struct {
	BaseError
}

func NewValidationError(message string, errorCode int) *ValidationError {
	return &ValidationError{
		BaseError: BaseError{
			Message:        message,
			HTTPStatusCode: 400,
			ErrorCode:      errorCode,
		},
	}
}

// package utils

// import (
// 	"encoding/json"
// 	"net/http"
// 	"os"
//
//
//
//
//
// 
// 
// )

//
// type CustomError struct {
// 	Message        string `json:"message"`
// 	ErrorCode      int    `json:"errorCode,omitempty"`
// 	HTTPStatusCode int    `json:"httpStatusCode,omitempty"`
// 	Service        string `json:"service,omitempty"`
// 	Success        bool   `json:"success,omitempty"`
// }

// func (e *CustomError) Error() string {
// 	return e.Message
// }

// var serviceName = os.Getenv("SERVICE_NAME")

// func NewUnauthorizedError(message string) *CustomError {
// 	return &CustomError{
// 		Message:        message,
// 		ErrorCode:      401,
// 		HTTPStatusCode: http.StatusUnauthorized,
// 		Service:        serviceName,
// 		Success:        false,
// 	}
// }

// func NewBadRequestError(message string) *CustomError {
// 	return &CustomError{
// 		Message:        message,
// 		ErrorCode:      400,
// 		HTTPStatusCode: http.StatusBadRequest,
// 		Service:        serviceName,
// 		Success:        false,
// 	}
// }

// func NewConflictError(message string) *CustomError {
// 	return &CustomError{
// 		Message:        message,
// 		ErrorCode:      409,
// 		HTTPStatusCode: http.StatusConflict,
// 		Service:        serviceName,
// 		Success:        false,
// 	}
// }

// func NewInternalServerError(message string) *CustomError {
// 	return &CustomError{
// 		Message:        message,
// 		ErrorCode:      500,
// 		HTTPStatusCode: http.StatusInternalServerError,
// 		Service:        serviceName,
// 		Success:        false,
// 	}
// }

// func NewUnauthenticatedError(message string) *CustomError {
// 	return &CustomError{
// 		Message:        message,
// 		ErrorCode:      401,
// 		HTTPStatusCode: http.StatusUnauthorized,
// 		Service:        serviceName,
// 		Success:        false,
// 	}
// }

// func NewNotFoundError(message string) *CustomError {
// 	return &CustomError{
// 		Message:        message,
// 		ErrorCode:      404,
// 		HTTPStatusCode: http.StatusNotFound,
// 		Service:        serviceName,
// 		Success:        false,
// 	}
// }

// func (e *CustomError) ToJSON() ([]byte, error) {
// 	return json.Marshal(e)
// }
