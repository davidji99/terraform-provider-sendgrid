---
layout: "sendgrid"
page_title: "Provider: Sendgrid"
sidebar_current: "docs-sendgrid-index"
description: |-
  Terraform provider Sendgrid.
---

# Sendgrid Provider

The Sendgrid provider is used to interact with resources provided by [Sendgrid API v3](https://sendgrid.com/docs/api-reference/)
and needs to be configured with credentials before it can be used.

## Contributing

Development happens in the [GitHub repo](https://github.com/davidji99/terraform-provider-sendgrid):

* [Releases](https://github.com/davidji99/terraform-provider-sendgrid/releases)
* [Issues](https://github.com/davidji99/terraform-provider-sendgrid/issues)

## Example Usage

```hcl
provider "sendgrid" {
  # ...
}

# Create a new API key.
resource "sendgrid_api_key" "admin" {
  # ...
}
```

## Authentication

The Sendgrid provider offers a flexible means of providing credentials for authentication.
The following methods are supported, listed in order of precedence, and explained below:

- Static credentials
- Environment variables

### Static credentials

Credentials can be provided statically by adding an `api_key` arguments to the Sendgrid provider block:

```hcl
provider "sendgrid" {
  api_key = "SOME_API_KEY"
}
```

### Environment variables

When the Sendgrid provider block does not contain an `api_key` argument, the missing credentials will be sourced
from the environment via the `SENDGRID_API_KEY` environment variables respectively:

```hcl
provider "sendgrid" {}
```

```shell
$ export SENDGRID_API_KEY="SOME_KEY"
$ terraform plan
Refreshing Terraform state in-memory prior to plan...
```

## Argument Reference

The following arguments are supported:

* `api_key` - (Required) Sendgrid API key. It must be provided, but it can also
  be sourced from [other locations](#Authentication).

* `base_url` - (Optional) Custom API URL.
  Can also be sourced from the `SENDGRID_API_URL` environment variable.
