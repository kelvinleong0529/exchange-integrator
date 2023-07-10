package interactive_broker_client_portal

import (
	"exchange-integrator/utils/network"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// DO NOT pass orderId when creating a new alert, toolID is only required for MTA alert
func (i *InteractiveBrokerClientPortalManager) CreateOrModifyAlert(queryParam *CreateOrModifyAlertReq) (*CreateOrModifyAlertRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf(endpointCreateOrModifyAlert, queryParam.AccountID),
	}
	result := new(CreateOrModifyAlertRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}
