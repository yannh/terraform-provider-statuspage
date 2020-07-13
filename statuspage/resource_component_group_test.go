package statuspage

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccStatuspageComponentGroup_Basic(t *testing.T) {
	rid := acctest.RandIntRange(1000, 9999)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComponentGroup_basic(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component_group.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component_group.default", "description", "Acc. Tests"),
					resource.TestCheckResourceAttr("statuspage_component_group.default", "components.#", "1"),
				),
			},
			{
				Config: testAccComponentGroup_update(rid),
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

	pageID = fmt.Sprintf("page-id-%d", rid)
	componentGroupID := rid

	resourceName := "statuspage_component_group.default"
	importInput := fmt.Sprintf("%s/%d", pageID, componentGroupID)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComponentGroup_basic(rid),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     importInput,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComponentGroup_basic(rand int) string {
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

func testAccComponentGroup_update(rand int) string {
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
