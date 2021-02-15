---
layout: "sendgrid"
page_title: "HerokuX: sendgrid_api_key"
sidebar_current: "docs-sendgrid-resource-api-key"
description: |-
Provides the ability to manage Sendgrid API keys.
---

# sendgrid\_api\_key

This resource provides the ability to manage [Sendgrid API keys](https://sendgrid.com/docs/ui/account-and-settings/api-keys).

-> **IMPORTANT!**
Please be very careful when deleting this resource as any deleted API keys are NOT recoverable and invalidated immediately.
Furthermore, this resource renders the `key` attribute in plain-text in your state file.
Please ensure that your state file is properly secured and encrypted at rest.

## Example Usage

```hcl-terraform
resource "sendgrid_api_key" "foobar" {
  name = "name_of_my_api_key"
  scopes = ["api_keys.read", "templates.create"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) `<string>` The name of your API Key.

* `scopes` - (Optional) `<string>` The permissions this API Key will have access to.
  Please note the following:

    * You may not give an API key greater permissions than the API key used to authenticate with the provider.
      For example, if the authenticating key has XYZ permissions, a new API key created from this resource can
      only have at most, XYZ permissions. It cannot have ABC permissions.

    * Not specifying this attributes results in the new API key having all permissions
      available to the authenticating key.

## Attributes Reference

The following attributes are exported:

* `key` - The actual API key. This attribute value does not get displayed in logs or regular output.

## Import

An existing API key can be imported using its key ID.

For example:

```shell script
$ terraform import sendgrid_api_key.foobar "<API_KEY_ID>"
```

-> **IMPORTANT!**
Due to API limitations, the actual API key is not shown in the response, so an imported `sendgrid_api_key`
will not have its `key` attribute set in state. Instead, `key` will be set to a placeholder value.