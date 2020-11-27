package error

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Methods(t *testing.T) {
	// System Error
	err := &SystemError{StatusCode: 0, Err: fmt.Errorf("Zero")}
	assert.Equal(t, "Zero", err.Err.Error())
	assert.Equal(t, 0, err.StatusCode)
	assert.Equal(t, "error [0]: Zero", err.Error())

	// Custom System Error
	err = CustomSystemError(1, "One")
	assert.Equal(t, "One", err.Err.Error())
	assert.Equal(t, 1, err.StatusCode)
	assert.Equal(t, "error [1]: One", err.Error())

	// Bad Request
	err = CustomBadRequest("Bad Request")
	assert.Equal(t, "Bad Request", err.Err.Error())
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "error [400]: Bad Request", err.Error())

	err = BadRequest()
	assert.Equal(t, "Bad Request", err.Err.Error())
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "error [400]: Bad Request", err.Error())

	// Unauthorized Access
	err = CustomUnauthorizedAccess("Unauthorized Access")
	assert.Equal(t, "Unauthorized Access", err.Err.Error())
	assert.Equal(t, http.StatusUnauthorized, err.StatusCode)
	assert.Equal(t, "error [401]: Unauthorized Access", err.Error())

	err = UnauthorizedAccess()
	assert.Equal(t, "Unauthorized Access", err.Err.Error())
	assert.Equal(t, http.StatusUnauthorized, err.StatusCode)
	assert.Equal(t, "error [401]: Unauthorized Access", err.Error())

	// Forbidden Access
	err = CustomForbiddenAccess("Forbidden Access")
	assert.Equal(t, "Forbidden Access", err.Err.Error())
	assert.Equal(t, http.StatusForbidden, err.StatusCode)
	assert.Equal(t, "error [403]: Forbidden Access", err.Error())

	err = ForbiddenAccess()
	assert.Equal(t, "Forbidden Access", err.Err.Error())
	assert.Equal(t, http.StatusForbidden, err.StatusCode)
	assert.Equal(t, "error [403]: Forbidden Access", err.Error())

	// Not Found
	err = CustomNotFound("Resource Not Found")
	assert.Equal(t, "Resource Not Found", err.Err.Error())
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, "error [404]: Resource Not Found", err.Error())

	err = NotFound()
	assert.Equal(t, "Resource Not Found", err.Err.Error())
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, "error [404]: Resource Not Found", err.Error())

	// Conflict
	err = CustomConflict("Resource Conflict")
	assert.Equal(t, "Resource Conflict", err.Err.Error())
	assert.Equal(t, http.StatusConflict, err.StatusCode)
	assert.Equal(t, "error [409]: Resource Conflict", err.Error())

	err = Conflict()
	assert.Equal(t, "Resource Conflict", err.Err.Error())
	assert.Equal(t, http.StatusConflict, err.StatusCode)
	assert.Equal(t, "error [409]: Resource Conflict", err.Error())

	// Internal System Error
	err = CustomInternalSystemError("Internal System Error")
	assert.Equal(t, "Internal System Error", err.Err.Error())
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error [500]: Internal System Error", err.Error())

	err = InternalSystemError()
	assert.Equal(t, "Internal System Error", err.Err.Error())
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error [500]: Internal System Error", err.Error())

}
