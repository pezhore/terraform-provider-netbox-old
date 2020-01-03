package netbox

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a schema.Provider for Netbox
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_API_KEY", nil),
				Description: "The API key from Netbox, you may retreive this from your profile in the Netbox UI",
			},

			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_URL", nil),
				Description: "The Netbox API Base URL",
			},
			"skip_tls_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_SKIP_TLS_VERIFY", false),
				Description: "Optionally ignore any TLS verification, defaults to false",
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"netbox_ip_address": resourceIPAddress(),
		},
	}

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIKey:        d.Get("api_key").(string),
		Endpoint:      d.Get("url").(string),
		SkipTLSVerify: d.Get("skip_tls_verify").(bool),
	}

	return config, nil
}
