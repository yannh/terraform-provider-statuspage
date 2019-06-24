package statuspage

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccStatuspageMetricProvider_Basic(t *testing.T) {
	rid := acctest.RandIntRange(1000, 9999)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricProvider_basic(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_metrics_provider.datadog", "id"),
					resource.TestCheckResourceAttr("statuspage_metrics_provider.datadog", "type", "Datadog"),
				),
			},
		},
	})
}

func testAccMetricProvider_basic(rand int) string {
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
	`, pageID, os.Getenv("DATADOG_API_KEY"), os.Getenv("DATADOG_APPLICATION_KEY"))
}
