package error

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Methods(t *testing.T) {
	err := ForbiddenAccess()

	assert.Equal(t, "Forbidden Access", err.Err.Error())
	assert.Equal(t, http.StatusForbidden, err.StatusCode)
	assert.Equal(t, "Status 403: Error => Forbidden Access", err.Error())
}
