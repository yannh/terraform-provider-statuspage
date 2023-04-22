![test](https://github.com/yannh/terraform-provider-statuspage/workflows/test/badge.svg) [![Go Report card](https://goreportcard.com/badge/github.com/yannh/terraform-provider-statuspage)](https://goreportcard.com/report/github.com/yannh/terraform-provider-statuspage)
<a href="https://terraform.io">
    <img src="https://raw.githubusercontent.com/hashicorp/terraform-website/master/public/img/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider for Statuspage.io

The Statuspage provider is used to interact with the resources supported by Statuspage.io.

 * [Documentation](https://registry.terraform.io/providers/yannh/statuspage/latest/docs)
 * [Download](https://github.com/yannh/terraform-provider-statuspage/releases)


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
$ export STATUSPAGE_PAGE=xxx    # PageID of the Statuspage page
$ export STATUSPAGE_PAGE_2=xxx  # PageID of another Statuspage page
$ export DATADOG_API_KEY=yyy
$ export DATADOG_APPLICATION_KEY=zzz
```

, and then run `make test acc`.

```sh
$ make test acc
```
