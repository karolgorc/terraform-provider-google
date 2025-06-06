// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/gemini/resource_gemini_code_tools_setting_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package gemini_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccGeminiCodeToolsSetting_geminiCodeToolsSettingBasicExample_update(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"setting_id": fmt.Sprintf("tf-test-ls-%s", acctest.RandString(t, 10)),
	}
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGeminiCodeToolsSetting_geminiCodeToolsSettingBasicExample_basic(context),
			},
			{
				ResourceName:            "google_gemini_code_tools_setting.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "code_tools_setting_id", "terraform_labels"},
			},
			{
				Config: testAccGeminiCodeToolsSetting_geminiCodeToolsSettingBasicExample_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_gemini_code_tools_setting.example", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_gemini_code_tools_setting.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "code_tools_setting_id", "terraform_labels"},
			},
		},
	})
}
func testAccGeminiCodeToolsSetting_geminiCodeToolsSettingBasicExample_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gemini_code_tools_setting" "example" {
    code_tools_setting_id = "%{setting_id}"
    location = "global"
    enabled_tool {
        handle = "my_handle"
        tool = "my_tool"
        account_connector = "my_con"
        config {
            key = "my_key"
            value = "my_value"
        }
        uri_override = "my_uri_override"
    }
}
`, context)
}
func testAccGeminiCodeToolsSetting_geminiCodeToolsSettingBasicExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_gemini_code_tools_setting" "example" {
    code_tools_setting_id = "%{setting_id}"
    location = "global"
    labels = {"my_key" = "my_value"}
    enabled_tool {
        handle = "my_handle"
        tool = "my_tool"
        account_connector = "my_con"
        config {
            key = "my_key"
            value = "my_value"
        }
        uri_override = "my_uri_override"
    }
}
`, context)
}
