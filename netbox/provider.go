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
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_API_KEY", ""),
				Description: descriptions["api_key"],
			},

			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_URL", nil),
				Description: descriptions["url"],
			},
			"skip_tls_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_SKIP_TLS_VERIFY", false),
				Description: descriptions["skip_tls_verify"],
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"netbox_ip_address": resourceASGroup(),
		},
	}
}

func init() {
	descriptions = map[string]string{
		"api_key": "The API key from Netbox, you may retreive this from your profile\n" +
			"in the Netbox UI",

		"url": "The Netbox API base URL.",

		"skip_tls_verify": "Optionally ignore TLS verification, defaults to false",
	}
}

func configureProvider(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := Config{
		APIKey:           d.Get("api_key").(string),
		Endpoint:         d.Get("url").(string),
		SkipTLSVerify:    d.Get("skip_tls_verify").(bool),
		terraformVersion: terraformVersion,
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}

	return config, nil
}
