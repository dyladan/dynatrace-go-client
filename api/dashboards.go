package api

import "gopkg.in/resty.v1"

type dashboardService service

func (s *dashboardService) GetAll() ([]Dashboard, *resty.Response, error) {

	dashboards := new(DashboardResponse)

	apiResponse, err := s.client.Do("GET", "/api/config/v1/dashboards", nil, dashboards)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return dashboards.Dashboards, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
