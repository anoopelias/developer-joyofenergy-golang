package usage

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

func TestGetUsage(t *testing.T) {
	accounts := repository.NewAccounts(map[string]string{"home-sweet-home": "test-plan"})
	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{"home-sweet-home": {{
			Time:    time.Now().Add(-24 * time.Hour),
			Reading: 5.0,
		}, {
			Time:    time.Now().Add(-24 * 2 * time.Hour),
			Reading: 15.0,
		}}},
	)
	pricePlans := repository.NewPricePlans(
		[]domain.PricePlan{{
			PlanName: "test-plan",
			UnitRate: 3.0,
		}},
		&meterReadings,
	)
	service := NewService(
		logrus.NewEntry(logrus.StandardLogger()),
		&meterReadings,
		&pricePlans,
		&accounts,
	)
	usage, err := service.GetUsage("home-sweet-home")
	expected := domain.Usage{
		SmartMeterId:      "home-sweet-home",
		LastWeekUsageCost: 10 * 24 * 7 * 3.0,
	}

	assert.NoError(t, err)
	assert.Equal(t, expected.SmartMeterId, usage.SmartMeterId)
	assert.Equal(t, expected.LastWeekUsageCost, usage.LastWeekUsageCost)
}

func TestGetUsage_IgnoreOlderReadings(t *testing.T) {
	accounts := repository.NewAccounts(map[string]string{"home-sweet-home": "test-plan"})
	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{"home-sweet-home": {{
			Time:    time.Now().Add(-24 * time.Hour),
			Reading: 5.0,
		}, {
			Time:    time.Now().Add(-24 * 2 * time.Hour),
			Reading: 15.0,
		}, {
			Time:    time.Now().Add(-24 * 10 * time.Hour),
			Reading: 200.0,
		}}},
	)
	pricePlans := repository.NewPricePlans(
		[]domain.PricePlan{{
			PlanName: "test-plan",
			UnitRate: 3.0,
		}},
		&meterReadings,
	)
	service := NewService(
		logrus.NewEntry(logrus.StandardLogger()),
		&meterReadings,
		&pricePlans,
		&accounts,
	)
	usage, err := service.GetUsage("home-sweet-home")
	expected := domain.Usage{
		SmartMeterId:      "home-sweet-home",
		LastWeekUsageCost: 10 * 24 * 7 * 3.0,
	}

	assert.NoError(t, err)
	assert.Equal(t, expected.SmartMeterId, usage.SmartMeterId)
	assert.Equal(t, expected.LastWeekUsageCost, usage.LastWeekUsageCost)
}

func TestGetUsageError_NoPricePlan(t *testing.T) {
	accounts := repository.NewAccounts(map[string]string{})
	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{"home-sweet-home": {{
			Time:    time.Now().Add(-24 * time.Hour),
			Reading: 5.0,
		}, {
			Time:    time.Now().Add(-24 * 2 * time.Hour),
			Reading: 15.0,
		}}},
	)
	pricePlans := repository.NewPricePlans(
		[]domain.PricePlan{{
			PlanName: "test-plan123",
			UnitRate: 3.0,
		}},
		&meterReadings,
	)
	service := NewService(
		logrus.NewEntry(logrus.StandardLogger()),
		&meterReadings,
		&pricePlans,
		&accounts,
	)
	_, err := service.GetUsage("home-sweet-home")
	expected := domain.ErrNoPricePlan

	assert.EqualError(t, expected, err.Error())
}

func TestGetUsageError_UnknownPricePlan(t *testing.T) {
	accounts := repository.NewAccounts(map[string]string{"home-sweet-home": "test-plan"})
	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{"home-sweet-home": {{
			Time:    time.Now().Add(-24 * time.Hour),
			Reading: 5.0,
		}, {
			Time:    time.Now().Add(-24 * 2 * time.Hour),
			Reading: 15.0,
		}}},
	)
	pricePlans := repository.NewPricePlans(
		[]domain.PricePlan{{
			PlanName: "test-plan123",
			UnitRate: 3.0,
		}},
		&meterReadings,
	)
	service := NewService(
		logrus.NewEntry(logrus.StandardLogger()),
		&meterReadings,
		&pricePlans,
		&accounts,
	)
	_, err := service.GetUsage("home-sweet-home")
	expected := domain.ErrNoPricePlan

	assert.EqualError(t, expected, err.Error())
}
