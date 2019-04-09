package api

import (
	"fmt"
	"gopkg.in/resty.v1"
)

type dashboardService service

func (s *dashboardService) GetAll() ([]DashboardStub, *resty.Response, error) {

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

func (s *dashboardService) Get(ID string) (*Dashboard, *resty.Response, error) {

	dashboard := new(Dashboard)
	path := fmt.Sprintf("/api/config/v1/dashboards/%s", ID)

	apiResponse, err := s.client.Do("GET", path, nil, dashboard)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return dashboard, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
