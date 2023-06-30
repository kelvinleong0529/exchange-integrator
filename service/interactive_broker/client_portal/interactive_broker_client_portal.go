package interactive_broker_client_portal

import "net/http"

func NewInteractiveBrokerClientPortalManager() *InteractiveBrokerClientPortalManager {
	return new(InteractiveBrokerClientPortalManager)
}

func (i *InteractiveBrokerClientPortalRequest) NewHttpRequest() (*http.Request, error) {
	endpoint := i.endpoint

	request, err := http.NewRequest(i.method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return request, nil
}
