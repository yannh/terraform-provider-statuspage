package main

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yannh/terraform-provider-statuspage/go-statuspage"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"statuspage_component": resourceComponent(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var token string
	if v := os.Getenv("STATUSPAGE_TOKEN"); v != "" {
		token = v
	}

	return statuspage.NewClient(token), nil
}
