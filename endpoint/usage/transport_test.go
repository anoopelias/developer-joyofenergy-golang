package usage

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMakeGetUsageHandler(t *testing.T) {
	mockLogger := logrus.New().WithField("test", "mock")
	h := MakeGetUsageHandler(mockLogger)
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
