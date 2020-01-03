Terraform Netbox Provider
===================================

- Website: https://www.terraform.io

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12+
-	[Go](https://golang.org/doc/install) 1.12+ (to build the provider plugin)

Acknowledgements
----------------

This is a learning effort and is built with inspiration from other terraform providers, including:

*   [Huawei Cloud Stack Provider](https://github.com/huaweicloud/terraform-provider-huaweicloudstack)
*   [Limberger's Netbox Provider](https://github.com/limberger/terraform-provider-netbox)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-netbox`

```sh
$ git clone https://github.com/pezhore/terraform-provider-netbox $GOPATH/src/github.com/terraform-providers/terraform-provider-netbox
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-netbox
$ make build
```

## Quick Start

Username/Password Configuration

```hcl
# Configure the Netbox Provider with Username/Password 
# This will work with a single defined/default network, otherwise you need to specify network
# to fix errrors about multiple networks found.
provider "netbox" {
  url  = "http://netboxdocker.localhost"
  api_key = "0123456789abcdef0123456789abcdef01234567"
}
```

Using the provider
----------------------
Please see the documentation at [provider usage](website/docs/index.html.markdown).

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
 $GOPATH/bin/terraform-provider-netbox
...
```

## License

Terraform-Provider-Netbox is under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

