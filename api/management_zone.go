package api

import (
	resty "gopkg.in/resty.v1"
)

type managementZoneService service

// Gets all management zones that have been defined
func (m *managementZoneService) GetAll() ([]ManagementZoneType, *resty.Response, error) {
	managementZones := new(ManagementZoneResponse)

	apiResponse, err := m.client.Do("GET", "/api/config/v1/managementZones", nil, managementZones)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return managementZones.Values, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())
}

// Create creates a new management zone
func (m *managementZoneService) Create(managementZone ManagementZoneType) (*ManagementZoneResponse, *resty.Response, error) {
	managementZoneResponse := new(ManagementZoneResponse)

	apiResponse, err := m.client.Do("POST", "/api/config/v1/managementZones", managementZone, managementZoneResponse)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return managementZoneResponse, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())
}
