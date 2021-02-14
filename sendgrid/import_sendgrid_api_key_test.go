package sendgrid

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccSendgridApiKey_importBasic_withScopes(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))
	scopes := `"api_keys.read"`

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSendgridApiKey_basic_withScopes(name, scopes),
			},
			{
				ResourceName:            "sendgrid_api_key.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"key"},
			},
		},
	})
}

func TestAccSendgridApiKey_importBasic_withNoScopes(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSendgridApiKey_basic_noScopes(name),
			},
			{
				ResourceName:            "sendgrid_api_key.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"key"},
			},
		},
	})
}
