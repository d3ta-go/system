package error

import "fmt"

// SystemError type
type SystemError struct {
	StatusCode int
	Err        error
}

// Error string
func (e *SystemError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.StatusCode, e.Err)
}

// Extensions error Extensions
func (e *SystemError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.StatusCode,
		"message": e.Err,
	}
}
