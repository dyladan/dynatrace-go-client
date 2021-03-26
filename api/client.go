package api

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type service struct {
	client *Client
}

type Client struct {
	common service

	AutoTags     *autoTagsService
	Dashboards   *dashboardService
	Events       *eventsService
	CustomDevice *customDeviceService
	Problem      *problemService

	RestyClient *resty.Client

	Log *log.Logger
}

type Config struct {
	APIKey    string
	BaseURL   string
	Debug     bool
	Retries   int
	RetryTime time.Duration
	Log       *log.Logger
}

// New returns a new Client for the specified apiKey.
func New(config Config) Client {
	r := resty.New()
	r.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	r.SetTimeout(30 * time.Second)
	//r.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))

	baseURL := config.BaseURL
	if baseURL == "" {
		panic("Base URL required")
	}
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = fmt.Sprintf("%s/", baseURL)
	}

	r.SetHeader("Authorization", fmt.Sprintf("Api-Token %s", config.APIKey))
	r.SetHostURL(baseURL)

	if config.Debug {
		r.SetDebug(true)
	}

	if config.Log == nil {
		config.Log = log.StandardLogger()
	}

	r.RetryCount = 1
	if config.Retries != 0 {
		r.RetryCount = config.Retries
	}

	if config.RetryTime == 0*time.Millisecond {
		config.RetryTime = 100 * time.Millisecond
	}
	r.RetryWaitTime = config.RetryTime

	r.AddRetryCondition(
		func(r *resty.Response, e error) bool {
			return r.StatusCode() >= http.StatusBadRequest
		},
	)

	c := Client{
		RestyClient: r,
		Log:         config.Log,
	}

	c.common.client = &c
	c.AutoTags = (*autoTagsService)(&c.common)
	c.Dashboards = (*dashboardService)(&c.common)
	c.Events = (*eventsService)(&c.common)
	c.CustomDevice = (*customDeviceService)(&c.common)
	c.Problem = (*problemService)(&c.common)

	return c
}

func (c *Client) Do(method string, path string, body interface{}, response interface{}) (*resty.Response, error) {
	r := c.RestyClient.R().
		SetError(&ErrorResponse{}).
		SetHeader("Content-Type", "application/json")

	if body != nil {
		r = r.SetBody(body)
	}

	if response != nil {
		r = r.SetResult(response)
	}

	c.Log.WithFields(log.Fields{"path": path, "method": method, "body": fmt.Sprintf("%+v", body)}).Debug("Making HTTP Request")
	apiResponse, err := r.Execute(method, path)
	c.Log.WithFields(log.Fields{"path": path, "method": method, "apiResponse": fmt.Sprintf("%+v", apiResponse)}).Debug("Received HTTP Response")

	return apiResponse, err
}

func StatusError(status int) error {
	return fmt.Errorf("Unexpected status code: %d\n", status)
}
