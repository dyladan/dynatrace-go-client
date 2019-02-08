package api

import "fmt"

func (c *Client) QueryApplicationNameDetectionRules() ([]NameDetectionRule, error) {
	resp := NameDetectionResponse{}

	err := c.Do("GET", "/api/config/v1/applicationDetectionRules", nil, &resp)
	return resp.Values, err
}

func (c *Client) GetApplicationNameDetectionRule(id string) (NameDetectionRuleDetail, error) {
	resp := NameDetectionRuleDetail{}

	err := c.Do("GET", fmt.Sprintf("/api/config/v1/applicationDetectionRules/%s", id), nil, &resp)

	return resp, err
}

func (c *Client) DeleteApplicationNameDetectionRule(id string) error {

	err := c.Do("DELETE", fmt.Sprintf("/api/config/v1/applicationDetectionRules/%s", id), nil, nil)

	return err
}

func (c *Client) CreateApplicationNameDetectionRule(body NameDetectionRuleDetail) (NameDetectionRule, error) {
	resp := NameDetectionRule{}

	err := c.Do("POST", "/api/config/v1/applicationDetectionRules", body, &resp)
	return resp, err
}

func (c *Client) UpdateApplicationNameDetectionRule(id string, body NameDetectionRuleDetail) error {

	return c.Do("PUT", fmt.Sprintf("/api/config/v1/applicationDetectionRules/%s", id), body, nil)
}
