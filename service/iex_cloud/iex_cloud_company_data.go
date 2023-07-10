package iex_cloud

import (
	"exchange-integrator/utils/network"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Returns the latest compensation information for the CEO of the company matching the symbol.
func (i *IEXCloudManager) CEOCompensation(queryParam *CEOCompensationReq) (*CEOCompensationRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	request := &IEXCloudRequest{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf(endpointCEOCompensation, queryParam.Symbol),
	}
	result := new(CEOCompensationRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}
