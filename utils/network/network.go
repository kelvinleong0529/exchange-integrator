package network

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Request interface {
	NewHttpRequest() (*http.Request, error)
}

func SubmitRequestAndUnmarshalResponse(request Request, responseStruct interface{}) error {
	httpRequest, err := request.NewHttpRequest()
	if err != nil {
		return err
	}

	client := new(http.Client)

	response, err := client.Do(httpRequest)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, responseStruct)
}
