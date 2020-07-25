---
layout: "statuspage"
page_title: "Provider: Statuspage"
sidebar_current: "docs-statuspage-index"
description: |-
  Terraform provider for Statuspage.io
---

# Statuspage Provider

The Statuspage provider is used to interact with the resources supported by Statuspage.io.

Use the navigation to the left to read about the available resources.

## Example Usage

Authentication currently works by setting the environment variable `STATUSPAGE_TOKEN`, or by configuring the provider:

!> **Warning:** Hard-coding credentials into any Terraform configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

```hcl
provider "statuspage" {
  token = "YOURTOKEN"
}

# Example resource configuration
resource "statuspage_component" "my_component" {
     page_id     = "pageid"
     name        = "My Website"
     description = "Status of my website"
     status      = "operational"
 }
```
