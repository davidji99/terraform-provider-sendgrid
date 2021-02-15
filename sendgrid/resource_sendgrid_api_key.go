package sendgrid

import (
	"context"
	"fmt"
	"github.com/davidji99/terraform-provider-sendgrid/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceSendgridApiKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSendgridApiKeyCreate,
		ReadContext:   resourceSendgridApiKeyRead,
		UpdateContext: resourceSendgridApiKeyUpdate,
		DeleteContext: resourceSendgridApiKeyDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceSendgridApiKeyImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"scopes": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceSendgridApiKeyImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*Config).API

	key, _, getErr := client.APIKeys.Get(d.Id())
	if getErr != nil {
		return nil, fmt.Errorf("unable to retrieve API Key %s", d.Id())
	}

	d.SetId(key.GetID())
	d.Set("name", key.GetName())
	d.Set("key", "not available as resource was imported")

	setScopesState(d, key)

	return []*schema.ResourceData{d}, nil
}

func resourceSendgridApiKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API
	opts := &api.KeyRequest{}

	if v, ok := d.GetOk("name"); ok {
		opts.Name = v.(string)
		log.Printf("[DEBUG] new api_key name is : %v", opts.Name)
	}

	if v, ok := d.GetOk("scopes"); ok {
		var scopes []string
		vl := v.(*schema.Set).List()

		for _, l := range vl {
			scopes = append(scopes, l.(string))
		}

		opts.Scopes = scopes
		log.Printf("[DEBUG] new api_key scopes are : %v", opts.Scopes)
	}

	log.Printf("[DEBUG] Creating API Key named %s", opts.Name)

	key, _, createErr := client.APIKeys.Create(opts)
	if createErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Unable to create API Key %s", opts.Name),
			Detail:   createErr.Error(),
		})
		return diags
	}

	log.Printf("[DEBUG] Created API Key named %s", opts.Name)

	d.SetId(key.GetID())

	// Set the actual API key here as this value is not returned for a GET request.
	d.Set("key", key.GetKey())

	return resourceSendgridApiKeyRead(ctx, d, meta)
}

func resourceSendgridApiKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API

	key, _, getErr := client.APIKeys.Get(d.Id())
	if getErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Unable to retrieve API Key ID %s", d.Id()),
			Detail:   getErr.Error(),
		})
		return diags
	}

	d.Set("name", key.GetName())

	setScopesState(d, key)

	return diags
}

func setScopesState(d *schema.ResourceData, key *api.Key) {
	if key.HasScopes() {
		// TODO: Remove this super hack. For some reason, the "2fa_required" & "sender_verification_eligible" scopes
		// TODO: are added to user specified scopes even if they weren't specified in the configuration.
		// TODO: So to address this, the resource will 'remove' these two scopes from the existing key's list of scopes.
		scopesModified := make([]string, 0)

		for _, scope := range key.Scopes {
			if scope != "2fa_required" && scope != "sender_verification_eligible" {
				scopesModified = append(scopesModified, scope)
			}
		}

		d.Set("scopes", scopesModified)
	} else {
		d.Set("scopes", []string{})
	}
}

func resourceSendgridApiKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API

	// Define variables to track what attributes are changing for the plan.
	// This is important as it'll determine which endpoint to use.
	var nameChanged, scopeChanged bool
	var newName string
	var newScopes []string

	if d.HasChange("name") {
		nameChanged = true
		newName = d.Get("name").(string)
	}

	if d.HasChange("scopes") {
		scopeChanged = true

		for _, scope := range d.Get("scopes").(*schema.Set).List() {
			newScopes = append(newScopes, scope.(string))
		}
	}

	log.Printf("[DEBUG] Updating API Key")

	if nameChanged && scopeChanged {
		_, _, updateErr := client.APIKeys.UpdateNameScopes(d.Id(), newName, newScopes)
		if updateErr != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update API Key name & scopes",
				Detail:   updateErr.Error(),
			})
			return diags
		}
	}

	if nameChanged && !scopeChanged {
		_, _, updateErr := client.APIKeys.UpdateName(d.Id(), newName)
		if updateErr != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update API Key name",
				Detail:   updateErr.Error(),
			})
			return diags
		}
	}

	if !nameChanged && scopeChanged {
		_, _, updateErr := client.APIKeys.UpdateNameScopes(d.Id(), d.Get("name").(string), newScopes)
		if updateErr != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update API Key scopes",
				Detail:   updateErr.Error(),
			})
			return diags
		}
	}

	log.Printf("[DEBUG] Updated API Key")

	return resourceSendgridApiKeyRead(ctx, d, meta)
}

func resourceSendgridApiKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API

	_, deleteErr := client.APIKeys.Delete(d.Id())
	if deleteErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Unable to delete API Key ID %s", d.Id()),
			Detail:   deleteErr.Error(),
		})
		return diags
	}

	d.SetId("")

	return nil
}
