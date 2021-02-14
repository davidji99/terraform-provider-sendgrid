package api

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"time"
)

const (
	// DefaultAPIv3BaseURL is the base url for Sendgrid API v3.
	DefaultAPIv3BaseURL = "https://api.sendgrid.com/v3"

	// DefaultUserAgent is the user agent used when making API calls.
	DefaultUserAgent = "sendgrid-go"

	// DefaultContentTypeHeader is the default and Content-Type header.
	DefaultContentTypeHeader = "application/json"
)

// Client manages communication with Sendgrid APIs.
type Client struct {
	// HTTP client used to communicate with the API.
	http *simpleresty.Client

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// config represents all of the API's configurations.
	config *Config

	// Services used for talking to different parts of the Sendgrid APIv3.
	APIKeys *KeysService
}

// service represents the API service client.
type service struct {
	client *Client
}

func New(opts ...Option) (*Client, error) {
	config := &Config{
		APIv3BaseURL:      DefaultAPIv3BaseURL,
		UserAgent:         DefaultUserAgent,
		ContentTypeHeader: DefaultContentTypeHeader,
		AcceptHeader:      "application/json",
		APIKey:            "",
	}

	// Define any user custom Client settings
	if optErr := config.ParseOptions(opts...); optErr != nil {
		return nil, optErr
	}

	client := &Client{
		config: config,
		http:   simpleresty.NewWithBaseURL(config.APIv3BaseURL),
	}

	// Set headers
	client.setHeaders()

	// Inject services
	client.injectServices()

	return client, nil
}

// injectServices adds the services to the client.
func (c *Client) injectServices() {
	c.common.client = c
	c.APIKeys = (*KeysService)(&c.common)
}

func (c *Client) setHeaders() {
	c.http.SetHeader("Content-type", c.config.ContentTypeHeader).
		SetHeader("Accept", c.config.AcceptHeader).
		SetHeader("User-Agent", c.config.UserAgent).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.config.APIKey)).
		SetTimeout(2 * time.Minute).
		SetAllowGetMethodPayload(true)

	// Set additional headers
	if c.config.CustomHTTPHeaders != nil {
		c.http.SetHeaders(c.config.CustomHTTPHeaders)
	}
}
