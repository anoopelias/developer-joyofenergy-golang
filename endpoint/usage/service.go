package usage

import (
	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

type Service interface {
	GetUsage(smartMeterId string) (domain.Usage, error)
}

type service struct {
	logger        *logrus.Entry
	meterReadings *repository.MeterReadings
	pricePlans    *repository.PricePlans
	accounts      *repository.Accounts
}

func NewService(
	logger *logrus.Entry,
	meterReadings *repository.MeterReadings,
	pricePlans *repository.PricePlans,
	accounts *repository.Accounts,

) Service {
	return &service{
		logger:        logger,
		meterReadings: meterReadings,
		pricePlans:    pricePlans,
		accounts:      accounts,
	}
}

func (s *service) GetUsage(smartMeterId string) (domain.Usage, error) {
	avg := calculateAverageReading(s.meterReadings.GetReadings(smartMeterId))
	units := avg * 24 * 7

	plan, err := s.accounts.PricePlanIdForSmartMeterId(smartMeterId)
	if err != nil {
		return domain.Usage{}, err
	}

	unitCost, err := s.pricePlans.UnitCostForPricePlan(plan)
	if err != nil {
		return domain.Usage{}, err
	}

	cost := units * unitCost

	return domain.Usage{
		SmartMeterId:      smartMeterId,
		LastWeekUsageCost: cost,
	}, nil
}

func calculateAverageReading(electricityReadings []domain.ElectricityReading) float64 {
	sum := 0.0
	for _, r := range electricityReadings {
		sum += r.Reading
	}
	return sum / float64(len(electricityReadings))
}
