package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
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

func (s *problemService) Close(problemID string, comment string) (*resty.Response, error) {

	path := fmt.Sprintf("/api/v2/problems/%s/close", problemID)
	pc := ProblemClose{Message: comment}
	apiResponse, err := s.client.Do("POST", path, pc, nil)
	if err != nil {
		return apiResponse, err
	}
	if apiResponse.StatusCode()/100 == 2 {
		return apiResponse, nil
	}
	return apiResponse, StatusError(apiResponse.StatusCode())

}

func (s *problemService) ListV1(
	relativeTime string,
	startTimestamp int64,
	endTimestamp int64,
	status string,
	impactLevel string,
	severityLevel string,
	tags []string,
	expandDetails bool) ([]ProblemV1, *resty.Response, error) {

	problems := new(ProblemV1Result)

	var queryParameters []string
	if relativeTime != "" {
		queryParameters = append(queryParameters, fmt.Sprintf("problemSelector=%s", relativeTime))
	}
	if startTimestamp != 0 {
		queryParameters = append(queryParameters, fmt.Sprintf("startTimestamp=%d", startTimestamp))
	}
	if endTimestamp != 0 {
		queryParameters = append(queryParameters, fmt.Sprintf("endTimestamp=%d", endTimestamp))
	}
	if status != "" {
		queryParameters = append(queryParameters, fmt.Sprintf("status=%s", status))
	}
	if impactLevel != "" {
		queryParameters = append(queryParameters, fmt.Sprintf("impactLevel=%s", status))
	}
	if severityLevel != "" {
		queryParameters = append(queryParameters, fmt.Sprintf("severityLevel=%s", status))
	}
	if expandDetails != false {
		queryParameters = append(queryParameters, fmt.Sprintf("expandDetails=%t", expandDetails))
	}
	for _, tag := range tags {
		queryParameters = append(queryParameters, fmt.Sprintf("tag=%s", tag))
	}

	queryString := ""
	if len(queryParameters) != 0 {
		queryString = fmt.Sprintf("?%s", strings.Join(queryParameters, "&"))
	}
	path := fmt.Sprintf("/api/v1/problem/feed%s", queryString)

	apiResponse, err := s.client.Do("GET", path, nil, problems)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return problems.Result.Problems, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
