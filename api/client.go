package api

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

type service struct {
	client *Client
}

type Client struct {
	common service

	AutoTags *AutoTagsService

	RestyClient *resty.Client
}

type Config struct {
	APIKey  string
	BaseURL string
	Debug   bool
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

	c.common.client = &c
	c.AutoTags = (*AutoTagsService)(&c.common)

	return c
}

func (c *Client) Do(method string, path string, body interface{}, response interface{}) (*resty.Response, error) {
	r := c.RestyClient.R().SetHeader("Content-Type", "application/json")

	if body != nil {
		r = r.SetBody(body)
	}

	if response != nil {
		r = r.SetResult(response)
	}

	apiResponse, err := r.Execute(method, path)

	return apiResponse, err
}

func StatusError(status int) error {
	return fmt.Errorf("Unexpected status code: %d\n", status)
}
