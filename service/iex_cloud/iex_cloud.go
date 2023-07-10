package iex_cloud

import "net/http"

func NewIEXCloudManager(apiToken string) *IEXCloudManager {
	return &IEXCloudManager{
		apiToken: apiToken,
	}
}

func (i *IEXCloudRequest) NewHttpRequest() (*http.Request, error) {
	endpoint := i.endpoint
	return http.NewRequest(i.method, endpoint, nil)
}
