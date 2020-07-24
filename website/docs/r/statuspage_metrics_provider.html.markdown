---
layout: "statuspage"
page_title: "Statuspage: statuspage_metrics_provider"
sidebar_current: "docs-statuspage-metrics-provider"
description: |-
  Statuspage metrics provider in the Terraform provider Statuspage.
---

# statuspage_metrics_provider

Statuspage metrics provider in the Terraform provider Statuspage.

## Example Usage

```hcl
resource "statuspage_metrics_provider" "statuspage_pingdom" {
    page_id         = "pageid"
    email           = "myemail@provider.com"
    password        = "pingdom_password"
    application_key = "pingdomAppKey"
    type            = "Pingdom"
}
```


## Argument Reference

The following arguments are supported:

 * page_id - (Required) the id of the page this component belongs to
 * type - (Required) One of "Pingdom", "NewRelic", "Librato", "Datadog", or "Self"
 * email - Required by the Librato and Pingdom type metrics providers.
 * password - Required by the Pingdom-type metrics provider.
 * api_key - Required by the Datadog and NewRelic type metrics providers.
 * api_token - Required by the Librato type metrics provider.
 * application_key - Required by the Pingdom and Datadog type metrics providers.

## Import Example Usage

```sh
terraform import statuspage_metrics_provider.statuspage_pingdom my-page-id/my-metrics-provider-id
```