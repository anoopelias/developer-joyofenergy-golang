package repository

import (
	"joi-energy-golang/domain"
)

type Accounts struct {
	smartMeterToPricePlanAccounts map[string]string
}

func NewAccounts(smartMeterToPricePlanAccounts map[string]string) Accounts {
	return Accounts{
		smartMeterToPricePlanAccounts: smartMeterToPricePlanAccounts,
	}
}

func (a *Accounts) PricePlanIdForSmartMeterId(smartMeterId string) (string, error) {
	// TODO indicate missing value
	pp, ok := a.smartMeterToPricePlanAccounts[smartMeterId]
	if !ok {
		return "", domain.ErrNoPricePlan
	}
	return pp, nil
}
