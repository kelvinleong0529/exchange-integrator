package interactive_broker_client_portal

import (
	"exchange-integrator/utils/network"
	"net/http"
)

// If the gateway has not received any requests for several minutes an open session will automatically timeout.
// The tickle endpoint pings the server to prevent the session from ending.
func (i *InteractiveBrokerClientPortalManager) KeepAliveSession() error {
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodPost,
		endpoint: endpointKeepaliveSession,
	}
	return network.SubmitRequestAndUnmarshalResponse(request, new(struct{}))
}

// Logs the user out of the gateway session. Any further activity requires re-authentication.
func (i *InteractiveBrokerClientPortalManager) TerminateSession() (*TerminateSessionRes, error) {
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodPost,
		endpoint: endpointTerminateSession,
	}
	result := new(TerminateSessionRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}

// Validates the current session for the SSO user
func (i *InteractiveBrokerClientPortalManager) ValidateSSO() (*ValidateSSORes, error) {
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodGet,
		endpoint: endpointValidateSSO,
	}
	result := new(ValidateSSORes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}

// Current Authentication status to the Brokerage system.
// Market Data and Trading is not possible if not authenticated, e.g. authenticated shows false
func (i *InteractiveBrokerClientPortalManager) AuthenticationStatus() (*AuthenticationStatusRes, error) {
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodPost,
		endpoint: endpointAuthenticationStatus,
	}
	result := new(AuthenticationStatusRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}

// When using the CP Gateway, this endpoint provides a way to reauthenticate
// to the Brokerage system as long as there is a valid SSO session, see /sso/validate.
func (i *InteractiveBrokerClientPortalManager) ReAuthenticate() (*AuthenticationStatusRes, error) {
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodPost,
		endpoint: endpointReauthenticate,
	}
	result := new(AuthenticationStatusRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}
