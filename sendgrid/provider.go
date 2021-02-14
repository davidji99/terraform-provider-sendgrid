package sendgrid

import (
	"context"
	"github.com/davidji99/terraform-provider-sendgrid/api"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENDGRID_API_KEY", nil),
			},

			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENDGRID_API_URL", api.DefaultAPIv3BaseURL),
			},

			"headers": {
				Type:     schema.TypeMap,
				Elem:     schema.TypeString,
				Optional: true,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"sendgrid_api_key": resourceSendgridApiKey(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	log.Println("[INFO] Initializing Sendgrid Provider")

	var diags diag.Diagnostics

	config := NewConfig()

	if applySchemaErr := config.applySchema(d); applySchemaErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to retrieve and set provider attributes",
			Detail:   applySchemaErr.Error(),
		})

		return nil, diags
	}

	if token, ok := d.GetOk("api_key"); ok {
		config.apiKey = token.(string)
	}

	if err := config.initializeAPI(); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to initialize API client",
			Detail:   err.Error(),
		})

		return nil, diags
	}

	log.Printf("[DEBUG] Sendgrid Provider initialized")

	return config, diags
}
