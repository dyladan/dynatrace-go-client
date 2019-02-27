package api

import "gopkg.in/resty.v1"

type AutoTagsService service

func (s *AutoTagsService) GetAll() ([]AutoTag, *resty.Response, error) {

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
