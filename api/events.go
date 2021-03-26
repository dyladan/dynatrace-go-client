package api

import (
	"github.com/go-resty/resty/v2"
)

type eventsService service

// Create posts a new event
func (e *eventsService) Create(event EventCreation) (*EventStoreResult, *resty.Response, error) {
	eventStoreResult := new(EventStoreResult)

	apiResponse, err := e.client.Do("POST", "/api/v1/events", event, eventStoreResult)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return eventStoreResult, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
