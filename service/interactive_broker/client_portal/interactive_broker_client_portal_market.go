package interactive_broker_client_portal

import (
	"exchange-integrator/utils/network"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Using a direct connection to the market data farm, will provide a list of historical market data for given contract ID
func (i *InteractiveBrokerClientPortalManager) MarketDataHistory(queryParam *MarketDataHistoryReq) (*MarketDataHistoryRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodGet,
		endpoint: endpointMarketDataHistory,
	}
	result := new(MarketDataHistoryRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}
