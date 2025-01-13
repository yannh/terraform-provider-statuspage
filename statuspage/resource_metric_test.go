package statuspage

import (
	"fmt"
	"os"
	"testing"
	//	"github.com/hashicorp/terraform/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccStatuspageMetricBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_metrics_provider.datadog", "id"),
					resource.TestCheckResourceAttr("statuspage_metrics_provider.datadog", "type", "Datadog"),
				),
			},
		},
	})
}

func testAccMetricBasic() string {
	return fmt.Sprintf(`
	variable "pageid" {
		default = "%s"
	}

	variable "api_key" {
		default = "%s"
	}

	variable "application_key" {
		default = "%s"
	}

	resource "statuspage_metrics_provider" "datadog" {
		page_id = "${var.pageid}"
		api_key = "${var.api_key}"
		application_key = "${var.application_key}"
		type = "Datadog"
		metric_base_uri = "https://app.datadoghq.com/api/v1"
	}

	resource "statuspage_metric" "datadog_metric" {
		page_id             = "${var.pageid}"
		metrics_provider_id = "${statuspage_metrics_provider.datadog.id}"
		name                = "DataDog Agent Started"
		metric_identifier   = "avg:datadog.agent.started{*}.as_count()"
	}
	`, pageID, os.Getenv("DATADOG_API_KEY"), os.Getenv("DATADOG_APPLICATION_KEY"))
}
