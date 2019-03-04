package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
)

type autoTagsService service

// GetAll lists all configured auto tags
func (s *autoTagsService) GetAll() ([]AutoTag, *resty.Response, error) {

	autoTags := new(AutoTagResponse)

	apiResponse, err := s.client.Do("GET", "/api/config/v1/autoTags", nil, autoTags)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTags.Values, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// Create creates a new auto tag
func (s *autoTagsService) Create(autoTag AutoTag) (*AutoTag, *resty.Response, error) {
	autoTagResp := new(AutoTag)

	apiResponse, err := s.client.Do("POST", "/api/config/v1/autoTags", autoTag, autoTagResp)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTagResp, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// Get gets the properties of the specified auto tag
func (s *autoTagsService) Get(ID string, includeProcessGroupReferences bool) (*AutoTag, *resty.Response, error) {

	autoTag := new(AutoTag)

	path := fmt.Sprintf("/api/config/v1/autoTags/%s?includeProcessGroupReferences=%t", ID, includeProcessGroupReferences)

	apiResponse, err := s.client.Do("GET", path, nil, autoTag)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTag, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// Update updates an existing auto tag or creates a new one
func (s *autoTagsService) Update(ID string, autoTag AutoTag) (*AutoTag, *resty.Response, error) {
	autoTagResp := new(AutoTag)

	url := fmt.Sprintf("/api/config/v1/autoTags/%s", ID)
	apiResponse, err := s.client.Do("PUT", url, autoTag, autoTagResp)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode() == 204 {
		return nil, apiResponse, nil
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTagResp, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// Delete deletes the specified auto tag
func (s *autoTagsService) Delete(ID string) (*resty.Response, error) {

	url := fmt.Sprintf("/api/config/v1/autoTags/%s", ID)
	apiResponse, err := s.client.Do("DELETE", url, nil, nil)

	if err != nil {
		return nil, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return apiResponse, nil
	}

	return apiResponse, StatusError(apiResponse.StatusCode())

}

// ValidateUpdate validates update of existing auto tags for the `PUT /autoTags/{id}` request
func (s *autoTagsService) ValidateUpdate(ID string, autoTag AutoTag) (*Error, *resty.Response, error) {

	validatorResp := new(Error)
	url := fmt.Sprintf("/api/config/v1/autoTags/%s/validator", ID)

	apiResponse, err := s.client.Do("POST", url, autoTag, nil)

	if apiResponse.StatusCode() == 400 {

		unmarshalError := json.Unmarshal(apiResponse.Body(), validatorResp)
		if unmarshalError != nil {
			return nil, apiResponse, unmarshalError
		}
		return validatorResp, apiResponse, nil
	}

	if apiResponse.StatusCode()/100 == 2 {
		return nil, apiResponse, nil
	}

	if err != nil {
		return nil, nil, err
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// ValidateCreate validates new auto tags for the `POST /autoTags` request
func (s *autoTagsService) ValidateCreate(autoTag AutoTag) (*Error, *resty.Response, error) {

	validatorResp := new(Error)

	apiResponse, err := s.client.Do("POST", "/api/config/v1/autoTags/", autoTag, nil)

	if apiResponse.StatusCode() == 400 {

		unmarshalError := json.Unmarshal(apiResponse.Body(), validatorResp)
		if unmarshalError != nil {
			return nil, apiResponse, unmarshalError
		}
		return validatorResp, apiResponse, nil
	}

	if apiResponse.StatusCode()/100 == 2 {
		return nil, apiResponse, nil
	}

	if err != nil {
		return nil, nil, err
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
