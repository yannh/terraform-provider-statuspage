[![pipeline status](https://gitlab.com/yannhamon/terraform-provider-statuspage/badges/master/pipeline.svg)](https://gitlab.com/yannhamon/terraform-provider-statuspage/commits/master)

# Terraform Provider for Statuspage.io

The Statuspage provider is used to interact with the resources supported by Statuspage.io.
Authentication currently works by setting the environment variable STATUSPAGE_TOKEN or by configuring the provider:

```
provider "statuspage" {
  token = "YOURTOKEN"
}
```

## Download

You can download the [latest build from Gitlab](https://gitlab.com/yannhamon/terraform-provider-statuspage/-/jobs/artifacts/master/download?job=build)


## statuspage_component

Components are the individual pieces of infrastructure that are listed on your status page.

### Example usage

```
resource "statuspage_component" "my_component" {
    page_id     = "pageid"
    name        = "My Website"
    description = "Status of my website"
    status      = "operational"
}
```

### Argument Reference

The following arguments are supported:

 * page_id - (Required) the id of the page this component belongs to
 * name - (Required) Name of the component
 * description - Description of the component
 * status - status of the component - must be one of "operational", "under_maintenance", "degraded_performance", "partial_outage", "major_outage" or ""
 * only_show_if_degraded (bool) - Should this component be shown component only if in degraded state
 * showcase (bool) - Should this component be showcased

The following attributes are exported:

 * automation_email - Email address to send automation events to

### Import Example Usage

```sh
terraform import statuspage_component.my_component my-page-id/my-component-id
```

## statuspage_component_group

Component groups provide a way to organize components. When a group is deleted, its child components will be orphaned. Note: A group cannot be empty, so if all the child components are deleted, the group will be deleted automatically. Another implication of this is that components must be created before their groups, when a group is created it will require a list of component IDs.

### Example usage

```
resource "statuspage_component_group" "my_group" {
    page_id     = "pageid"
    name        = "terraform"
    description = "Created by terraform"
    components  = ["${statuspage_component.my_component.id}"]
}
```

### Argument Reference

The following arguments are supported:

 * page_id - (Required) the id of the page this component belongs to
 * components - (Required) List of component IDs
 * name - (Required) name of the component group
 * description - description of the component group

### Import Example Usage

```sh
terraform import statuspage_component_group.my_group my-page-id/my-component-group-id
```

## statuspage_metrics

System metrics are a great way to build trust and transparency around your organization, and ensure that your page is doing work for you each and every day.

### Example usage

```
resource "statuspage_metric" "website_metrics" {
    page_id             = "pageid"
    metrics_provider_id = "${statuspage_metrics_provider.statuspage_pingdom.id}"
    name                = "My Website"
    metric_identifier   = "pingdom_check_id"
}
```

### Argument Reference

The following arguments are supported:

 * page_id - (Required) the id of the page this component belongs to
 * name - Name of metric
 * metric_identifier - The identifier used to look up the metric data from the provider
 * transform - The transform to apply to metric before pulling into Statuspage. One of: "average", "count", "max", "min", or "sum"
 * suffix - Suffix to describe the units on the graph
 * y_axis_min - The lower bound of the y axis
 * y_axis_max - The upper bound of the y axis
 * y_axis_hidden - Should the values on the y axis be hidden on render
 * display - Should the metric be displayed
 * decimal_places - How many decimal places to render on the graph
 * tooltip_description - A description for the tooltip

### Import Example Usage

```sh
terraform import statuspage_metric.website_metrics my-page-id/my-metric-id
```

## statuspage_metrics_provider

### Example usage

```
resource "statuspage_metrics_provider" "statuspage_pingdom" {
    page_id         = "pageid"
    email           = "myemail@provider.com"
    password        = "pingdom_password"
    application_key = "pingdomAppKey"
    type            = "Pingdom"
}
```

### Argument Reference

The following arguments are supported:

 * page_id - (Required) the id of the page this component belongs to
 * type - (Required) One of "Pingdom", "NewRelic", "Librato", "Datadog", or "Self"
 * email - Required by the Librato and Pingdom type metrics providers.
 * password - Required by the Pingdom-type metrics provider.
 * api_key - Required by the Datadog and NewRelic type metrics providers.
 * api_token - Required by the Librato type metrics provider.
 * application_key - Required by the Pingdom and Datadog type metrics providers.

### Import Example Usage

```sh
terraform import statuspage_metrics_provider.statuspage_pingdom my-page-id/my-metrics-provider-id
```

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.10.x
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)


## Building The Provider

Clone repository to: `$GOPATH/src/github.com/yannh/terraform-provider-statuspage

```sh
$ mkdir -p $GOPATH/src/github.com/yannh; cd $GOPATH/src/github.com/yannh
$ git clone https://github.com/yannh/terraform-provider-statuspage.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/yannh/terraform-provider-statuspage
$ make build
```


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-statuspage
...
```

In order to test the provider, set the following variables:

```sh
$ export STATUSPAGE_TOKEN=www
$ export STATUSPAGE_PAGE=xxx
$ export DATADOG_API_KEY=yyy
$ export DATADOG_APPLICATION_KEY=zzz
```

, and then run `make test acc`.

```sh
$ make test acc
```
