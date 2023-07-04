package interactive_broker_client_portal

import (
	"exchange-integrator/utils/network"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// In non-tiered account structures, returns a list of accounts for which the user can view position and account information.
// This endpoint must be called prior to calling other /portfolio endpoints for those accounts. For querying a list of accounts
// which the user can trade, see /iserver/accounts. For a list of subaccounts in tiered account structures (e.g. financial
// advisor or ibroker accounts) see /portfolio/subaccounts.
func (i *InteractiveBrokerClientPortalManager) PortfolioAccounts() (*PortfolioAccountsRes, error) {
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodGet,
		endpoint: endpointPortfolioAccounts,
	}
	result := new(PortfolioAccountsRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}

// Account information related to account Id /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.
func (i *InteractiveBrokerClientPortalManager) AccountInformation(queryParam *AccountInformationReq) (*PortfolioAccountsRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf(endpointAccountInformation, queryParam.AccountID),
	}
	result := new(PortfolioAccountsRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}

// If an user has multiple accounts, and user wants to get orders, trades, etc. of an account other than currently selected
// account, then user can update the currently selected account using this API and then can fetch required information for
// the newly updated account.
func (i *InteractiveBrokerClientPortalManager) SwitchAccount(queryParam *SwitchAccountReq) (*SwitchAccountRes, error) {
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodPost,
		endpoint: endpointSwitchAccount,
	}
	result := new(SwitchAccountRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}
