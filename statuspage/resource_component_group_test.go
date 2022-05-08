package statuspage

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccStatuspageComponentGroupBasic(t *testing.T) {
	rid := acctest.RandIntRange(1000, 9999)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComponentGroupBasic(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component_group.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component_group.default", "description", "Acc. Tests"),
					resource.TestCheckResourceAttr("statuspage_component_group.default", "components.#", "1"),
				),
			},
			{
				Config: testAccComponentGroupUpdate(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component_group.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component_group.default", "description", "Acc. Tests"),
					resource.TestCheckResourceAttr("statuspage_component_group.default", "components.#", "4"),
				),
			},
		},
	})
}

func TestAccStatuspageComponentGroup_BasicImport(t *testing.T) {
	rid := acctest.RandIntRange(1000, 9999)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComponentGroupBasic(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component_group.default", "id"),
				),
			},
			{
				ResourceName:      "statuspage_component_group.default",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(ts *terraform.State) (string, error) {
					rs := ts.RootModule().Resources["statuspage_component_group.default"]
					return fmt.Sprintf("%s/%s", pageID, rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccComponentGroupBasic(rand int) string {
	return fmt.Sprintf(`
	variable "component_name" {
		default = "tf-testacc-component-group-%d"
	}

	variable "pageid" {
		default = "%s"
	}

	resource "statuspage_component" "comp1" {
		page_id = "${var.pageid}"
		name = "${var.component_name}_component"
		description = "test component"
		status = "operational"
	}

	resource "statuspage_component_group" "default" {
		page_id     = "${var.pageid}"
		name        = "${var.component_name}"
		description = "Acc. Tests"
		components  = ["${statuspage_component.comp1.id}"]
	}
	`, rand, pageID)
}

func testAccComponentGroupUpdate(rand int) string {
	return fmt.Sprintf(`
	variable "component_name" {
		default = "tf-testacc-component-group-%d"
	}

	variable "pageid" {
		default = "%s"
	}

	resource "statuspage_component" "comp1" {
		page_id = "${var.pageid}"
		name = "${var.component_name}_component"
		description = "test component"
		status = "operational"
	}

	resource "statuspage_component" "comp2" {
		page_id = "${var.pageid}"
		name = "${var.component_name}_component"
		description = "test component"
		status = "operational"
	}

	resource "statuspage_component" "comp3" {
		page_id = "${var.pageid}"
		name = "${var.component_name}_component"
		description = "test component"
		status = "operational"
	}

	resource "statuspage_component" "comp4" {
		page_id = "${var.pageid}"
		name = "${var.component_name}_component"
		description = "test component"
		status = "operational"
	}

	resource "statuspage_component_group" "default" {
		page_id     = "${var.pageid}"
		name        = "${var.component_name}"
		description = "Acc. Tests"
		components  = ["${statuspage_component.comp1.id}", "${statuspage_component.comp2.id}", "${statuspage_component.comp3.id}", "${statuspage_component.comp4.id}"]
	}
	`, rand, pageID)
}
