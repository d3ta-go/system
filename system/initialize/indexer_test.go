package initialize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenAllIndexerConnection(t *testing.T) {
	h, err := newHandler(t)
	if assert.NoError(t, err, "Error while creating handler: newHandler") {
		if !assert.NotNil(t, h) {
			return
		}
	}

}
