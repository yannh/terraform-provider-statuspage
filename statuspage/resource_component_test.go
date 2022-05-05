package statuspage

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccStatuspageComponentBasic(t *testing.T) {
	rid := acctest.RandIntRange(1000, 9999)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComponentBasic(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component.default", "description", "test component"),
					resource.TestCheckResourceAttr("statuspage_component.default", "group_id", ""),
					resource.TestCheckResourceAttr("statuspage_component.default", "status", "operational"),
					resource.TestCheckResourceAttr("statuspage_component.default", "showcase", "true"),
				),
			},
			{
				Config: testAccComponentUpdate(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component.default", "description", "updated component"),
					resource.TestCheckResourceAttr("statuspage_component.default", "group_id", ""),
					resource.TestCheckResourceAttr("statuspage_component.default", "status", "major_outage"),
					resource.TestCheckResourceAttr("statuspage_component.default", "showcase", "false"),
				),
			},
		},
	})
}

func TestAccStatuspageComponentBasicPageIDUpdate(t *testing.T) {
	rid := acctest.RandIntRange(1000, 9999)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComponentBasic(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component.default", "page_id", pageID),
					resource.TestCheckResourceAttr("statuspage_component.default", "description", "test component"),
					resource.TestCheckResourceAttr("statuspage_component.default", "group_id", ""),
					resource.TestCheckResourceAttr("statuspage_component.default", "status", "operational"),
					resource.TestCheckResourceAttr("statuspage_component.default", "showcase", "true"),
				),
			},
			{
				Config: testAccComponentUpdatePageID(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component.default", "page_id", pageID2),
					resource.TestCheckResourceAttr("statuspage_component.default", "description", "updated component"),
					resource.TestCheckResourceAttr("statuspage_component.default", "status", "major_outage"),
					resource.TestCheckResourceAttr("statuspage_component.default", "group_id", ""),
					resource.TestCheckResourceAttr("statuspage_component.default", "showcase", "false"),
				),
			},
			{
				Config: testAccComponentUpdateGroupID(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component.comp1", "id"),
					resource.TestCheckResourceAttr("statuspage_component.comp1", "group_id", ""),
					resource.TestCheckResourceAttrSet("statuspage_component_group.default", "id"),
					resource.TestCheckResourceAttrSet("statuspage_component.comp2", "id"),
					resource.TestCheckResourceAttrSet("statuspage_component.comp2", "group_id"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccStatuspageComponent_BasicImport(t *testing.T) {
	rid := acctest.RandIntRange(1000, 9999)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComponentBasic(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component.default", "description", "test component"),
				),
			},
			{
				ResourceName:      "statuspage_component.default",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(ts *terraform.State) (string, error) {
					rs := ts.RootModule().Resources["statuspage_component.default"]
					return fmt.Sprintf("%s/%s", pageID, rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccComponentBasic(rand int) string {
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
		showcase = true
	}
	`, rand, pageID)
}

func testAccComponentUpdate(rand int) string {
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
		description = "updated component"
		status = "major_outage"
		showcase = false
	}
	`, rand, pageID)
}

func testAccComponentUpdatePageID(rand int) string {
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
		description = "updated component"
		status = "major_outage"
		showcase = false
	}
	`, rand, pageID2)
}

func testAccComponentUpdateGroupID(rand int) string {
	return fmt.Sprintf(`
	variable "name" {
		default = "tf-testacc-component-%d"
	}

	variable "pageid" {
		default = "%s"
	}

	resource "statuspage_component" "comp1" {
		page_id = "${var.pageid}"
		name = "${var.name}-extra"
		description = "component to facilitate group creation"
		status = "operational"
		showcase = false
	}

	resource "statuspage_component_group" "default" {
		page_id     = "${var.pageid}"
		name        = "${var.name}"
		description = "Acc. Tests"
		components  = ["${statuspage_component.comp1.id}"]
	}

	resource "statuspage_component" "comp2" {
		page_id = "${var.pageid}"
		name = "${var.name}"
		group_id = "${statuspage_component_group.default.id}"
		description = "updated component"
		status = "major_outage"
		showcase = false
	}
	`, rand, pageID2)
}
