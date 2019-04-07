package statuspage

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	sp "github.com/yannh/statuspage-go-sdk"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("STATUSPAGE_TOKEN", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"statuspage_component":        resourceComponent(),
			"statuspage_component_group":  resourceComponentGroup(),
			"statuspage_metric":           resourceMetric(),
			"statuspage_metrics_provider": resourceMetricsProvider(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return sp.NewClient(d.Get("token").(string)), nil
}
