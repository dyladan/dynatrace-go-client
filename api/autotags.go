package api

import (
	"fmt"
	"gopkg.in/resty.v1"
)

type autoTagsService service

func (s *autoTagsService) GetAll() ([]AutoTag, *resty.Response, error) {

	autoTags := new(AutoTagResponse)

	apiResponse, err := s.client.Do("GET", "/api/config/v1/autoTags", nil, autoTags)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTags.Values, apiResponse, err
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

func (s *autoTagsService) Create(autoTag AutoTag) (*AutoTag, *resty.Response, error) {
	autoTagResp := new(AutoTag)

	apiResponse, err := s.client.Do("POST", "/api/config/v1/autoTags", autoTag, autoTagResp)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTagResp, apiResponse, err
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

func (s *autoTagsService) Get(ID string, includeProcessGroupReferences bool) (*AutoTag, *resty.Response, error) {

	autoTag := new(AutoTag)

	path := fmt.Sprintf("/api/config/v1/autoTags/%s?includeProcessGroupReferences=%t", ID, includeProcessGroupReferences)

	apiResponse, err := s.client.Do("GET", path, nil, autoTag)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTag, apiResponse, err
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
