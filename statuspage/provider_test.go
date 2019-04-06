package statuspage

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"statuspage": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("STATUSPAGE_TOKEN"); v == "" {
		t.Error("Environment variable STATUSPAGE_TOKEN needs to be set for acceptance testing")
	}

	if v := os.Getenv("STATUSPAGE_PAGE"); v == "" {
		t.Error("Environment variable STATUSPAGE_PAGE needs to be set for acceptance testing")
	}
}
