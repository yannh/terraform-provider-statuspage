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
    api_token       = "a-pingdom-api-token"
    type            = "Pingdom"
}
```

## Argument Reference

The following arguments are supported:

 * `page_id` - (Required) the id of the page this component belongs to
 * `type` - (Required) One of `Pingdom`, `NewRelic`, `Librato`, `Datadog`, or `Self`
 * `email` - Required by the Librato and Pingdom type metrics providers.
 * `password` - Required by the Pingdom-type metrics provider.
 * `api_key` - Required by the Datadog and NewRelic type metrics providers.
 * `api_token` - Required by the Librato and Pingdom-type metrics provider.
 * `application_key` - Required by the Pingdom and Datadog type metrics providers.

## Import

`statuspage_metrics_provider` can be imported using the ID of the metrics provider, e.g.

```sh
$ terraform import statuspage_metrics_provider.statuspage_pingdom my-page-id/my-metrics-provider-id
```
