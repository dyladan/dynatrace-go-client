package api

import (
	"fmt"
	"gopkg.in/resty.v1"
	"strings"
)

type problemService service

func (s *problemService) List(fields []string, problemSelector string, entitySelector string, sort string) ([]Problem, *resty.Response, error) {

	problems := new(ProblemsResponse)

	var queryParameters []string
	if fields != nil {
		queryParameters = append(queryParameters, fmt.Sprintf("fields=%s", strings.Join(fields, ",")))
	}
	if problemSelector != "" {
		queryParameters = append(queryParameters, fmt.Sprintf("problemSelector=%s", problemSelector))
	}
	if entitySelector != "" {
		queryParameters = append(queryParameters, fmt.Sprintf("entitySelector=%s", entitySelector))
	}
	if sort != "" {
		queryParameters = append(queryParameters, fmt.Sprintf("sort=%s", sort))
	}

	queryString := ""
	if len(queryParameters) != 0 {
		queryString = fmt.Sprintf("?%s", strings.Join(queryParameters, "&"))
	}
	path := fmt.Sprintf("/api/v2/problems%s", queryString)

	apiResponse, err := s.client.Do("GET", path, nil, problems)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return problems.Problems, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
