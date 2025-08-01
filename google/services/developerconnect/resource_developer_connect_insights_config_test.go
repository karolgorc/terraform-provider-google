// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/developerconnect/resource_developer_connect_insights_config_test.go
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package developerconnect_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDeveloperConnectInsightsConfig_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
	}

	acctest.SkipIfVcr(t) // See: https://github.com/GoogleCloudPlatform/magic-modules/pull/14412
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDeveloperConnectInsightsConfig_basic(context),
			},
			{
				ResourceName:            "google_developer_connect_insights_config.insights_config",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"insights_config_id", "labels", "location", "terraform_labels", "workload"},
			},
			{
				Config: testAccDeveloperConnectInsightsConfig_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_developer_connect_insights_config.insights_config", plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
			{
				ResourceName:            "google_developer_connect_insights_config.insights_config",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"insights_config_id", "location", "labels", "terraform_labels", "workload"},
			},
		},
	})
}

func testAccDeveloperConnectInsightsConfig_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
	resource "google_project" "project" {
		project_id = "dci-tf-%{random_suffix}"
		name = "Service Project"
		org_id          = "%{org_id}"
  		billing_account = "%{billing_account}"
		deletion_policy = "DELETE"
	}
	
	# Grant Permissions
	resource "google_project_iam_member" "apphub_permissions" {
		project = google_project.project.project_id
		role = "roles/apphub.admin"
		member = "serviceAccount:hashicorp-test-runner@ci-test-project-188019.iam.gserviceaccount.com"
	}

	resource "google_project_iam_member" "insights_agent" {
		project = google_project.project.project_id
		role = "roles/developerconnect.insightsAgent"
		member = "serviceAccount:66214305248-compute@developer.gserviceaccount.com"
	}

	# Enable APIs
	resource "google_project_service" "apphub_api_service" {
		project = google_project.project.project_id
		service = "apphub.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "containeranalysis_api" {
		project = google_project.project.project_id
		service = "containeranalysis.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "containerscanning_api" {
		project = google_project.project.project_id
		service = "containerscanning.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "container_api" {
		project = google_project.project.project_id
		service = "container.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "artifactregistry_api" {
		project = google_project.project.project_id
		service = "artifactregistry.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "cloudbuild_api" {
		project = google_project.project.project_id
		service = "cloudbuild.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "cloudasset_api" {
		project = google_project.project.project_id
		service = "cloudasset.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "compute_api" {
		project = google_project.project.project_id
		service = "compute.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "devconnect_api" {
		project = google_project.project.project_id
		service = "developerconnect.googleapis.com"
		depends_on = [google_project.project]
	}

	# Wait delay after enabling APIs and granting permissions
	resource "time_sleep" "wait_for_propagation" {
		depends_on = [
			google_project_iam_member.apphub_permissions,
			google_project_iam_member.insights_agent,
			google_project_service.apphub_api_service,
			google_project_service.containeranalysis_api,
			google_project_service.containerscanning_api,
			google_project_service.container_api,
			google_project_service.artifactregistry_api,
			google_project_service.artifactregistry_api,
			google_project_service.cloudbuild_api,
			google_project_service.cloudasset_api,
			google_project_service.compute_api,
			google_project_service.devconnect_api,
		]
		create_duration  = "120s"
	}

	resource "google_apphub_application" "my_apphub_application" {
		location = "us-central1"
		application_id = "tf-test-example-application%{random_suffix}"
		scope {
			type = "REGIONAL"
		}
		project = google_project.project.project_id
		depends_on = [time_sleep.wait_for_propagation]
	}
	
	resource "google_developer_connect_insights_config" "insights_config" {
		location           = "us-central1"
		insights_config_id = "tf-test-ic%{random_suffix}"
		project            = google_project.project.project_id
		annotations = {}
    	labels = {}
    	app_hub_application = format("//apphub.googleapis.com/projects/%s/locations/%s/applications/%s",
			google_project.project.number,
        	google_apphub_application.my_apphub_application.location,
        	google_apphub_application.my_apphub_application.application_id)
		
		depends_on = [time_sleep.wait_for_propagation]
    }
  `, context)
}

func testAccDeveloperConnectInsightsConfig_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
	resource "google_project" "project" {
		project_id = "dci-tf-%{random_suffix}"
		name = "Service Project"
		org_id          = "%{org_id}"
  		billing_account = "%{billing_account}"
		deletion_policy = "DELETE"
	}
	
	# Grant Permissions
	resource "google_project_iam_member" "apphub_permissions" {
		project = google_project.project.project_id
		role = "roles/apphub.admin"
		member = "serviceAccount:hashicorp-test-runner@ci-test-project-188019.iam.gserviceaccount.com"
	}

	resource "google_project_iam_member" "insights_agent" {
		project = google_project.project.project_id
		role = "roles/developerconnect.insightsAgent"
		member = "serviceAccount:66214305248-compute@developer.gserviceaccount.com"
	}

	# Enable APIs
		resource "google_project_service" "apphub_api_service" {
		project = google_project.project.project_id
		service = "apphub.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "containeranalysis_api" {
		project = google_project.project.project_id
		service = "containeranalysis.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "containerscanning_api" {
		project = google_project.project.project_id
		service = "containerscanning.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "container_api" {
		project = google_project.project.project_id
		service = "container.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "artifactregistry_api" {
		project = google_project.project.project_id
		service = "artifactregistry.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "cloudbuild_api" {
		project = google_project.project.project_id
		service = "cloudbuild.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "cloudasset_api" {
		project = google_project.project.project_id
		service = "cloudasset.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "compute_api" {
		project = google_project.project.project_id
		service = "compute.googleapis.com"
		disable_dependent_services=true
		depends_on = [google_project.project]
	}

	resource "google_project_service" "devconnect_api" {
		project = google_project.project.project_id
		service = "developerconnect.googleapis.com"
		depends_on = [google_project.project]
	}

	# Wait delay after enabling APIs and granting permissions
	resource "time_sleep" "wait_for_propagation" {
		depends_on = [
			google_project_iam_member.apphub_permissions,
			google_project_iam_member.insights_agent,
			google_project_service.apphub_api_service,
			google_project_service.containeranalysis_api,
			google_project_service.containerscanning_api,
			google_project_service.container_api,
			google_project_service.artifactregistry_api,
			google_project_service.artifactregistry_api,
			google_project_service.cloudbuild_api,
			google_project_service.cloudasset_api,
			google_project_service.compute_api,
			google_project_service.devconnect_api,
		]
		create_duration  = "120s"
	}

	resource "google_apphub_application" "my_apphub_application" {
		location = "us-central1"
		application_id = "tf-test-example-application%{random_suffix}"
		scope {
			type = "REGIONAL"
		}
		project = google_project.project.project_id
		depends_on = [time_sleep.wait_for_propagation]
	}
	resource "google_developer_connect_insights_config" "insights_config" {
		location           = "us-central1"
		insights_config_id = "tf-test-ic%{random_suffix}"
		project            = google_project.project.project_id
		annotations = {}
    	labels = {}
    	app_hub_application = format("//apphub.googleapis.com/projects/%s/locations/%s/applications/%s",
        	google_project.project.number,
        	google_apphub_application.my_apphub_application.location,
        	google_apphub_application.my_apphub_application.application_id)
		artifact_configs {
			google_artifact_analysis {
				project_id = google_project.project.project_id
			}
			google_artifact_registry {
				artifact_registry_package = "my-package"
				project_id                = google_project.project.project_id
			}
			uri = "us-docker.pkg.dev/my-project/my-repo/my-image"
		}
		depends_on = [time_sleep.wait_for_propagation]
    }
  `, context)
}
