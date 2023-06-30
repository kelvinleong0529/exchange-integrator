package interactive_broker_third_party

import "github.com/go-playground/validator/v10"

func (i *InteractiveBrokerThirdPartyManager) FinancialInstruments(queryParam *FinancialInstrumentsReq) (*FinancialInstrumentsRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	return nil, nil
}
