Terraform Provider for Statuspage.io
====================================

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10.x
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Usage
-----

```
resource "statuspage_component" "my_component" {
    page_id = "pageid"
    name = "My Website"
    description = "Status of my website"
    status = "operational"
}
```

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/yannh/terraform-provider-statuspage

```sh
$ mkdir -p $GOPATH/src/github.com/yannh; cd $GOPATH/src/github.com/yannh
$ git clone git@github.com:yannh/terraform-provider-statuspage
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/yannh/terraform-provider-statuspage
$ make build
```


Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-statuspage
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```