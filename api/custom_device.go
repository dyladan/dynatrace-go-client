package api

import (
	"fmt"
	"gopkg.in/resty.v1"
)

type customDeviceService service

//
// Create posts a new custom device
func (e *customDeviceService) Create(id string, cd CustomDevicePushMessage) (*CustomDevicePushResult, *resty.Response, error) {
	customDevicePushResult := new(CustomDevicePushResult)

	uri := fmt.Sprintf("/api/v1/entity/infrastructure/custom/%s", id)
	apiResponse, err := e.client.Do("POST", uri, cd, customDevicePushResult)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return customDevicePushResult, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
