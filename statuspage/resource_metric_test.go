package statuspage

import (
	"fmt"
	"os"
	"testing"
	//	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccStatuspageMetricBasic(t *testing.T) {
	//	resource.Test(t, resource.TestCase{
	//		PreCheck:  func() { testAccPreCheck(t) },
	//		Providers: testAccProviders,
	//		Steps: []resource.TestStep{
	//			{
	//				Config: testAccMetric_basic(),
	//				Check: resource.ComposeTestCheckFunc(
	//					resource.TestCheckResourceAttrSet("statuspage_metrics_provider.datadog", "id"),
	//					resource.TestCheckResourceAttr("statuspage_metrics_provider.datadog", "type", "Datadog"),
	//				),
	//			},
	//		},
	//	})
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
	}

	resource "statuspage_metric" "datadog_metric" {
		page_id             = "${var.pageid}"
		metrics_provider_id = "${statuspage_metrics_provider.datadog.id}"
		name                = "Forum"
		metric_identifier   = "sum:apache.net.request_per_s{project:forum}"
	}
	`, pageID, os.Getenv("DATADOG_API_KEY"), os.Getenv("DATADOG_APPLICATION_KEY"))
}
