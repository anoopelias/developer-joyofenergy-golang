package usage

import (
	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
)

type Service interface {
	GetUsage(smartMeterId string) (domain.Usage, error)
}

type service struct {
	logger *logrus.Entry
}

func NewService(
	logger *logrus.Entry,
) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) GetUsage(smartMeterId string) (domain.Usage, error) {
	return domain.Usage{
		SmartMeterId: smartMeterId,
	}, nil
}
