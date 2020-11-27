package error

import (
	"fmt"
	"net/http"
)

// CustomSystemError custom SystemError
func CustomSystemError(code int, message string) *SystemError {
	return &SystemError{StatusCode: code, Err: fmt.Errorf(message)}
}

// CustomBadRequest custom bad request error
func CustomBadRequest(message string) *SystemError {
	return CustomSystemError(http.StatusBadRequest, message)
}

// BadRequest error
//	Similar with: 400 Http Status Code
func BadRequest() *SystemError {
	return CustomBadRequest("Bad Request")
}

// CustomUnauthorizedAccess error
func CustomUnauthorizedAccess(message string) *SystemError {
	return CustomSystemError(http.StatusUnauthorized, message)
}

// UnauthorizedAccess error
//	Similar with: 401 Http Status Code
func UnauthorizedAccess() *SystemError {
	return CustomUnauthorizedAccess("Unauthorized Access")
}

// CustomForbiddenAccess error
func CustomForbiddenAccess(message string) *SystemError {
	return &SystemError{StatusCode: http.StatusForbidden, Err: fmt.Errorf(message)}
}

// ForbiddenAccess error
//	Similar with: 403 Http Status Code
func ForbiddenAccess() *SystemError {
	return CustomForbiddenAccess("Forbidden Access")
}

// CustomNotFound error
func CustomNotFound(message string) *SystemError {
	return CustomSystemError(http.StatusNotFound, message)
}

// NotFound error
//	Similar withL 404 Http Status Code
func NotFound() *SystemError {
	return CustomNotFound("Resource Not Found")
}

// CustomConflict error
//	Similar withL 409 Http Status Code
func CustomConflict(message string) *SystemError {
	return CustomSystemError(http.StatusConflict, message)
}

// Conflict error
//	Similar withL 409 Http Status Code
func Conflict() *SystemError {
	return CustomConflict("Resource Conflict")
}

// CustomInternalSystemError error
func CustomInternalSystemError(message string) *SystemError {
	return CustomSystemError(http.StatusInternalServerError, message)
}

// InternalSystemError error
//	Similar with: 500 Http Status Code
func InternalSystemError() *SystemError {
	return CustomInternalSystemError("Internal System Error")
}
