package usage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsageReturnResultFromService(t *testing.T) {
	e := makeUsageEndpoint()

	response, err := e(context.Background(), "123")

	assert.NoError(t, err)
	assert.Equal(t, nil, response)
}
