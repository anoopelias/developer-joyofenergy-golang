package usage

import (
	"joi-energy-golang/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type MockService struct {
	err error
	Service
}

func (s *MockService) GetUsage(smartMeterId string) (domain.Usage, error) {
	return domain.Usage{}, s.err
}

func TestMakeGetUsageHandler(t *testing.T) {
	mockService := &MockService{}
	mockLogger := logrus.New().WithField("test", "mock")
	h := MakeGetUsageHandler(mockService, mockLogger)
	r := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/usage/smart-meter-12345", nil)
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	actualStatusCode := result.StatusCode
	assert.Equal(t, http.StatusOK, actualStatusCode)
	err := result.Body.Close()
	assert.NoError(t, err)
}
