package statuspage

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var pageID string
var pageID2 string

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"statuspage": testAccProvider,
	}
	pageID = os.Getenv("STATUSPAGE_PAGE")
	pageID2 = os.Getenv("STATUSPAGE_PAGE_2")

}

func testAccPreCheck(t *testing.T) {
	v := ""
	if v = os.Getenv("STATUSPAGE_TOKEN"); v == "" {
		t.Error("Environment variable STATUSPAGE_TOKEN needs to be set for acceptance testing")
	}

	if v = os.Getenv("STATUSPAGE_PAGE"); v == "" {
		t.Error("Environment variable STATUSPAGE_PAGE needs to be set for acceptance testing")
	}

	if v = os.Getenv("STATUSPAGE_PAGE_2"); v == "" {
		t.Error("Environment variable STATUSPAGE_PAGE_2 needs to be set for acceptance testing")
	}
}
