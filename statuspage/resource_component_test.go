package statuspage

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccStatuspageComponent_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComponent_basic(acctest.RandIntRange(1000, 9999)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component.default", "description", "test component"),
					resource.TestCheckResourceAttr("statuspage_component.default", "status", "operational"),
				),
			},
		},
	})
}

func testAccComponent_basic(rand int) string {
	return fmt.Sprintf(`
	variable "name" {
		default = "tf-testacc-component-%d"
	}

	variable "pageid" {
		default = "%s"
	}

	resource "statuspage_component" "default" {
		page_id = "${var.pageid}"
		name = "${var.name}"
		description = "test component"
		status = "operational"
	}
	`, rand, "8l7wwkjhvgg7")
}
