package sendgrid

import (
	"fmt"
	"github.com/davidji99/terraform-provider-sendgrid/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

type Config struct {
	API     *api.Client
	Headers map[string]string

	apiKey       string
	apiv3BaseURL string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) initializeAPI() error {
	api, clientInitErr := api.New(api.APIKey(c.apiKey), api.APIv3BaseURL(c.apiv3BaseURL))
	if clientInitErr != nil {
		return clientInitErr
	}

	c.API = api

	log.Printf("[INFO] Sendgrid Client configured")

	return nil
}

func (c *Config) applySchema(d *schema.ResourceData) (err error) {
	if v, ok := d.GetOk("headers"); ok {
		headersRaw := v.(map[string]interface{})
		h := make(map[string]string)

		for k, v := range headersRaw {
			h[k] = fmt.Sprintf("%v", v)
		}

		c.Headers = h
	}

	if v, ok := d.GetOk("base_url"); ok {
		vs := v.(string)
		c.apiv3BaseURL = vs
	}

	return nil
}
