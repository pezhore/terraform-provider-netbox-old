package netbox

import (
	"fmt"
)

func resourceIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPAddressCreate,
		Read:   resourceIPAddressRead,
		Update: resourceIPAddressUpdate,
		Delete: resourceIPAddressDelete,

		Schema: map[string]*schema.Schema{
			"input": &schema.Schema{
				Type:    schema.TypeString,
				Requred: true,
			},
		},
	}
}

func resourceIPAddressCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("input"))
	return resourceIPAddressRead(d, m)
}

func resourceIPAddressRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIPAddressUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceIPAddressRead(d, m)
}

func resourceIPAddressDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
