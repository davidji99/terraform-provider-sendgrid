package sendgrid

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccSendgridApiKey_BasicWithNoScopes(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSendgridApiKey_basic_noScopes(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", name),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "scopes.#"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
		},
	})
}

// This test only works with an authenticated API key that only has full scope to "API Keys".
func TestAccSendgridApiKey_BasicWithScopes(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))
	scopes := `"api_keys.read"`

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSendgridApiKey_basic_withScopes(name, scopes),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", name),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "scopes.#"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
		},
	})
}

func TestAccSendgridApiKey_UpdateJustNameNoScopesDefined(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))
	nameEdited := fmt.Sprintf("%s+edited", name)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSendgridApiKey_basic_noScopes(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", name),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "scopes.#"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
			{
				Config: testAccCheckSendgridApiKey_basic_noScopes(nameEdited),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", nameEdited),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "scopes.#"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
		},
	})
}

func TestAccSendgridApiKey_UpdateJustNameThenAddScopes(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))
	nameEdited := fmt.Sprintf("%s+edited", name)
	scopes := `"api_keys.read"`

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSendgridApiKey_basic_noScopes(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", name),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "scopes.#"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
			{
				Config: testAccCheckSendgridApiKey_basic_withScopes(nameEdited, scopes),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", nameEdited),
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "scopes.#", "1"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
		},
	})
}

func TestAccSendgridApiKey_UpdateNamesAndScopes(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))
	nameEdited := fmt.Sprintf("%s+edited", name)
	scopes := `"api_keys.read"`
	scopesUpdated := `"api_keys.read", "templates.create"`

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSendgridApiKey_basic_withScopes(name, scopes),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "scopes.#", "1"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
			{
				Config: testAccCheckSendgridApiKey_basic_withScopes(nameEdited, scopesUpdated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", nameEdited),
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "scopes.#", "2"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
		},
	})
}

func TestAccSendgridApiKey_JustScopes(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))
	scopes := `"api_keys.read"`
	scopesUpdated := `"api_keys.read", "templates.create"`

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSendgridApiKey_basic_withScopes(name, scopes),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "scopes.#", "1"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
			{
				Config: testAccCheckSendgridApiKey_basic_withScopes(name, scopesUpdated),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"sendgrid_api_key.foobar", "scopes.#", "2"),
					resource.TestCheckResourceAttrSet(
						"sendgrid_api_key.foobar", "key"),
				),
			},
		},
	})
}

func testAccCheckSendgridApiKey_basic_withScopes(name, scopes string) string {
	return fmt.Sprintf(`
resource "sendgrid_api_key" "foobar" {
	name = "%s"
	scopes = [%s]
}
`, name, scopes)
}

func testAccCheckSendgridApiKey_basic_noScopes(name string) string {
	return fmt.Sprintf(`
resource "sendgrid_api_key" "foobar" {
	name = "%s"
}
`, name)
}
