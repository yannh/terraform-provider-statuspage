---
layout: "statuspage"
page_title: "Statuspage: statuspage_component_group"
sidebar_current: "docs-statuspage-component-group"
description: |-
  Statuspage component_group in the Terraform provider Statuspage.
---

# statuspage_component_group

Statuspage component group in the Terraform provider Statuspage.

## Example Usage

```hcl
resource "statuspage_component_group" "my_group" {
    page_id     = "pageid"
    name        = "terraform"
    description = "Created by terraform"
    components  = ["${statuspage_component.my_component.id}"]
}
```

## Argument Reference

The following arguments are supported:

 * `page_id` - (Required) the id of the page this component belongs to
 * `components` - (Required) List of component IDs
 * `name` - (Required) name of the component group
 * `description` - description of the component group

## Import

`statuspage_component_group` can be imported using the ID of the component group, e.g.

```sh
$ terraform import statuspage_component_group.my_group my-page-id/my-component-group-id
```