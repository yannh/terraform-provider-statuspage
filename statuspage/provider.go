package statuspage

import (
	"log"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	sp "github.com/yannh/statuspage-go-sdk"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"statuspage_component":         resourceComponent(),
			"statuspage_component_group":   resourceComponentGroup(),
			"statuspage_metric":            resourceMetric(),
			"statuspage_metrics_provider":  resourceMetricsProvider(),
			"statuspage_incident_template": resourceIncidentTemplate(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var token string

	log.Printf("[INFO] Initializing Statuspage client\n")

	if v := os.Getenv("STATUSPAGE_TOKEN"); v != "" {
		token = v
	}

	return sp.NewClient(token), nil
}
