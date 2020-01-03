package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/pezhore/terraform-provider-netbox/netbox"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: netbox.Provider})
}
