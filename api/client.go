package api

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

type Client struct {
	RestyClient *resty.Client
}

type Config struct {
	APIKey  string
	BaseURL string
	Debug   bool
}

type ErrorResponse struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

func (e *ErrorResponse) Error() string {
	if e != nil && e.Message != nil {
		return *e.Message
	}
	return "Unknown error"
}

// New returns a new Client for the specified apiKey.
func New(config Config) Client {
	r := resty.New()

	baseURL := config.BaseURL
	if baseURL == "" {
		panic("Base URL required")
	}

	r.SetHeader("Authorization", fmt.Sprintf("Api-Token %s", config.APIKey))
	r.SetHostURL(baseURL)

	if config.Debug {
		r.SetDebug(true)
	}

	c := Client{
		RestyClient: r,
	}

	return c
}

func (c *Client) Do(method string, path string, body interface{}, response interface{}) error {
	r := c.RestyClient.R().SetError(&ErrorResponse{}).SetHeader("Content-Type", "application/json")

	if body != nil {
		r = r.SetBody(body)
	}

	if response != nil {
		r = r.SetResult(response)
	}

	apiResponse, err := r.Execute(method, path)

	if err != nil {
		return err
	}

	statusClass := apiResponse.StatusCode() / 100 % 10

	if statusClass == 2 {
		return nil
	}

	rawError := apiResponse.Error()

	if rawError != nil {
		apiError := rawError.(*ErrorResponse)

		if apiError.Message != nil {
			return nil
		}
	}

	return fmt.Errorf("Unexpected status %v returned from API", apiResponse.StatusCode())
}
