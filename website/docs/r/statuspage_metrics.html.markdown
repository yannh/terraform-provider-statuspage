---
layout: "statuspage"
page_title: "Statuspage: statuspage_metrics"
sidebar_current: "docs-statuspage-metrics"
description: |-
  Statuspage metrics in the Terraform provider Statuspage.
---

# statuspage_metrics

Statuspage metrics in the Terraform provider Statuspage.

## Example Usage

```hcl
resource "statuspage_metric" "website_metrics" {
    page_id             = "pageid"
    metrics_provider_id = "${statuspage_metrics_provider.statuspage_pingdom.id}"
    name                = "My Website"
    metric_identifier   = "pingdom_check_id"
}
```

## Argument Reference

The following arguments are supported:

 * `page_id` - (Required) the id of the page this component belongs to
 * `name` - Name of metric
 * `metric_identifier` - The identifier used to look up the metric data from the provider
 * `transform` - The transform to apply to metric before pulling into Statuspage. One of: `average`, `count`, `max`, `min`, `sum`, `response_time`, `uptime`. For pingdom metrics_provider allowed values are `response_time` and `uptime`.
 * `suffix` - Suffix to describe the units on the graph
 * `y_axis_min` - The lower bound of the y axis
 * `y_axis_max` - The upper bound of the y axis
 * `y_axis_hidden` - Should the values on the y axis be hidden on render
 * `display` - Should the metric be displayed
 * `decimal_places` - How many decimal places to render on the graph
 * `tooltip_description` - A description for the tooltip
