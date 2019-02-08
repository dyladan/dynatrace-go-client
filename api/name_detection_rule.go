package api

import "fmt"

// Get an application name detection rule. If the rule does not exist, an empty NameDetectionRuleDetail
// will be returned with an Id of ""
func (c *Client) GetApplicationNameDetectionRule(id string) (NameDetectionRuleDetail, error) {
	resp := NameDetectionRuleDetail{}

	apiResponse, err := c.Do("GET", fmt.Sprintf("/api/config/v1/applicationDetectionRules/%s", id), nil, &resp)

	if err != nil {
		return resp, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return resp, nil
	}

	if apiResponse.StatusCode() == 404 {
		return resp, nil
	}

	return resp, StatusError(apiResponse.StatusCode())
}

// Delete an application name detection rule.
func (c *Client) DeleteApplicationNameDetectionRule(id string) error {
	_, err := c.Do("DELETE", fmt.Sprintf("/api/config/v1/applicationDetectionRules/%s", id), nil, nil)

	return err
}

// Create an application name detection rule. If the API responds with a non-2xx status code, an error is returned.
func (c *Client) CreateApplicationNameDetectionRule(body NameDetectionRuleDetail) (NameDetectionRule, error) {
	resp := NameDetectionRule{}

	apiResponse, err := c.Do("POST", "/api/config/v1/applicationDetectionRules", body, &resp)

	if err != nil {
		return resp, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return resp, err
	}

	return resp, StatusError(apiResponse.StatusCode())
}

// Update an application name detection rule. If the API responds with a non-2xx status code, an error is returned.
func (c *Client) UpdateApplicationNameDetectionRule(id string, body NameDetectionRuleDetail) error {
	apiResponse, err := c.Do("PUT", fmt.Sprintf("/api/config/v1/applicationDetectionRules/%s", id), body, nil)

	if err != nil {
		return err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return nil
	}

	return StatusError(apiResponse.StatusCode())
}
