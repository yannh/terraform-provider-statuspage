---
layout: "statuspage"
page_title: "Statuspage: statuspage_component"
sidebar_current: "docs-statuspage-component"
description: |-
  Statuspage component in the Terraform provider Statuspage.
---

# statuspage_component

Statuspage component in the Terraform provider Statuspage.

## Example Usage

```hcl
resource "statuspage_component" "my_component" {
    page_id     = "pageid"
    name        = "My Website"
    description = "Status of my website"
    status      = "operational"
    group_id    = "Component group id"
}
```

To use the Statuspage console instead of Terraform to manage the status
of the component:

```hcl
resource "statuspage_component" "my_component" {
    page_id     = "pageid"
    name        = "My Website"
    description = "Status of my website"
    status      = "operational"
    group_id    = "Component group id"
    lifecycle {
        ignore_changes = [
            status,
        ]
    }
}
```

## Argument Reference

The following arguments are supported:

 * `page_id` - (Required) the id of the page this component belongs to
 * `name` - (Required) Name of the component
 * `description` - Description of the component
 * `status` - status of the component - must be one of `operational`, `under_maintenance`, `degraded_performance`, `partial_outage`, `major_outage` or ` `
 * `only_show_if_degraded` (bool) - Should this component be shown component only if in degraded state
 * `showcase` (bool) - Should this component be showcased
 * `group_id` (Optional) - (string) - Component Group ID

The following attributes are exported:

 * `automation_email` - Email address to send automation events to

## Import

`statuspage_component` can be imported using the ID of the component, e.g.

```sh
$ terraform import statuspage_component.my_component my-page-id/my-component-id
```

