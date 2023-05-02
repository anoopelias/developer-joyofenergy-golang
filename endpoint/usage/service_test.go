package usage

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

func TestGetUsage(t *testing.T) {
	service := NewService(
		logrus.NewEntry(logrus.StandardLogger()),
	)
	usage, err := service.GetUsage("home-sweet-home")
	expected := domain.Usage{
		SmartMeterId: "home-sweet-home",
	}

	assert.NoError(t, err)
	assert.Equal(t, expected.SmartMeterId, usage.SmartMeterId)
}
