package usage

import (
	"context"
	"joi-energy-golang/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsageReturnResultFromService(t *testing.T) {
	s := &MockService{}
	e := makeUsageEndpoint(s)

	response, err := e(context.Background(), "123")
	expectedResponse := domain.Usage{}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}
