// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package colab_test

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccColabSchedule_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"location":           envvar.GetTestRegionFromEnv(),
		"project_id":         envvar.GetTestProjectFromEnv(),
		"service_account":    envvar.GetTestServiceAccountFromEnv(t),
		"end_time":           time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, 10).Format(time.RFC3339),
		"start_time":         time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, 1).Format(time.RFC3339),
		"random_suffix":      acctest.RandString(t, 10),
		"updated_start_time": time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, 2).Format(time.RFC3339),
		"updated_end_time":   time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, 5).Format(time.RFC3339),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckColabScheduleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccColabSchedule_full(context),
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
			{
				Config: testAccColabSchedule_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_colab_schedule.schedule", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func TestAccColabSchedule_update_state(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"location":           envvar.GetTestRegionFromEnv(),
		"project_id":         envvar.GetTestProjectFromEnv(),
		"service_account":    envvar.GetTestServiceAccountFromEnv(t),
		"end_time":           time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, 10).Format(time.RFC3339),
		"start_time":         time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, 1).Format(time.RFC3339),
		"random_suffix":      acctest.RandString(t, 10),
		"updated_start_time": time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, 2).Format(time.RFC3339),
		"updated_end_time":   time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, 5).Format(time.RFC3339),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckColabScheduleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccColabSchedule_full(context),
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "desired_state"},
			},
			{
				Config: testAccColabSchedule_active(context),
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "desired_state"},
			},
			{
				Config: testAccColabSchedule_paused(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_colab_schedule.schedule", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "desired_state"},
			},
			{
				Config: testAccColabSchedule_active(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_colab_schedule.schedule", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "desired_state"},
			},
			{
				Config: testAccColabSchedule_full(context),
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "desired_state"},
			},
			{
				Config: testAccColabSchedule_paused(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_colab_schedule.schedule", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "desired_state"},
			},
			{
				Config: testAccColabSchedule_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_colab_schedule.schedule", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "desired_state"},
			},
			{
				Config: testAccColabSchedule_active(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_colab_schedule.schedule", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_colab_schedule.schedule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "desired_state"},
			},
		},
	})
}

func testAccColabSchedule_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "my_runtime_template" {
  name = "tf-test-runtime-template%{random_suffix}"
  display_name = "Runtime template"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}

resource "google_storage_bucket" "output_bucket" {
  name          = "tf_test_my_bucket%{random_suffix}"
  location      = "US"
  force_destroy = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "notebook" {
  name   = "hello_world.ipynb"
  bucket = google_storage_bucket.output_bucket.name
  content = <<EOF
    {
      "cells": [
        {
          "cell_type": "code",
          "execution_count": null,
          "metadata": {},
          "outputs": [],
          "source": [
            "print(\"Hello, World!\")"
          ]
        }
      ],
      "metadata": {
        "kernelspec": {
          "display_name": "Python 3",
          "language": "python",
          "name": "python3"
        },
        "language_info": {
          "codemirror_mode": {
            "name": "ipython",
            "version": 3
          },
          "file_extension": ".py",
          "mimetype": "text/x-python",
          "name": "python",
          "nbconvert_exporter": "python",
          "pygments_lexer": "ipython3",
          "version": "3.8.5"
        }
      },
      "nbformat": 4,
      "nbformat_minor": 4
    }
    EOF
}

resource "google_colab_schedule" "schedule" {
  display_name = "tf-test-schedule%{random_suffix}"
  location = "%{location}"
  allow_queueing = true
  max_concurrent_run_count = 2
  cron = "TZ=America/Los_Angeles * * * * *"
  max_run_count = 5
  start_time = "%{start_time}"
  end_time = "%{end_time}"

  create_notebook_execution_job_request {
    notebook_execution_job {
      display_name = "Notebook execution"
      gcs_notebook_source {
        uri = "gs://${google_storage_bucket_object.notebook.bucket}/${google_storage_bucket_object.notebook.name}"
        generation = google_storage_bucket_object.notebook.generation
      }

      notebook_runtime_template_resource_name = "projects/${google_colab_runtime_template.my_runtime_template.project}/locations/${google_colab_runtime_template.my_runtime_template.location}/notebookRuntimeTemplates/${google_colab_runtime_template.my_runtime_template.name}"
      gcs_output_uri = "gs://${google_storage_bucket.output_bucket.name}"
      service_account = "%{service_account}"
      }
  }

  depends_on = [
    google_colab_runtime_template.my_runtime_template,
    google_storage_bucket.output_bucket,
  ]
}
`, context)
}

func testAccColabSchedule_paused(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "my_runtime_template" {
  name = "tf-test-runtime-template%{random_suffix}"
  display_name = "Runtime template"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}

resource "google_storage_bucket" "output_bucket" {
  name          = "tf_test_my_bucket%{random_suffix}"
  location      = "US"
  force_destroy = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "notebook" {
  name   = "hello_world.ipynb"
  bucket = google_storage_bucket.output_bucket.name
  content = <<EOF
    {
      "cells": [
        {
          "cell_type": "code",
          "execution_count": null,
          "metadata": {},
          "outputs": [],
          "source": [
            "print(\"Hello, World!\")"
          ]
        }
      ],
      "metadata": {
        "kernelspec": {
          "display_name": "Python 3",
          "language": "python",
          "name": "python3"
        },
        "language_info": {
          "codemirror_mode": {
            "name": "ipython",
            "version": 3
          },
          "file_extension": ".py",
          "mimetype": "text/x-python",
          "name": "python",
          "nbconvert_exporter": "python",
          "pygments_lexer": "ipython3",
          "version": "3.8.5"
        }
      },
      "nbformat": 4,
      "nbformat_minor": 4
    }
    EOF
}

resource "google_colab_schedule" "schedule" {
  display_name = "tf-test-schedule%{random_suffix}"
  location = "%{location}"
  allow_queueing = true
  max_concurrent_run_count = 2
  cron = "TZ=America/Los_Angeles * * * * *"
  max_run_count = 5
  start_time = "%{start_time}"
  end_time = "%{end_time}"

  desired_state = "PAUSED"

  create_notebook_execution_job_request {
    notebook_execution_job {
      display_name = "Notebook execution"
      gcs_notebook_source {
        uri = "gs://${google_storage_bucket_object.notebook.bucket}/${google_storage_bucket_object.notebook.name}"
        generation = google_storage_bucket_object.notebook.generation
      }

      notebook_runtime_template_resource_name = "projects/${google_colab_runtime_template.my_runtime_template.project}/locations/${google_colab_runtime_template.my_runtime_template.location}/notebookRuntimeTemplates/${google_colab_runtime_template.my_runtime_template.name}"
      gcs_output_uri = "gs://${google_storage_bucket.output_bucket.name}"
      service_account = "%{service_account}"
      }
  }

  depends_on = [
    google_colab_runtime_template.my_runtime_template,
    google_storage_bucket.output_bucket,
  ]
}
`, context)
}

func testAccColabSchedule_active(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "my_runtime_template" {
  name = "tf-test-runtime-template%{random_suffix}"
  display_name = "Runtime template"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}

resource "google_storage_bucket" "output_bucket" {
  name          = "tf_test_my_bucket%{random_suffix}"
  location      = "US"
  force_destroy = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "notebook" {
  name   = "hello_world.ipynb"
  bucket = google_storage_bucket.output_bucket.name
  content = <<EOF
    {
      "cells": [
        {
          "cell_type": "code",
          "execution_count": null,
          "metadata": {},
          "outputs": [],
          "source": [
            "print(\"Hello, World!\")"
          ]
        }
      ],
      "metadata": {
        "kernelspec": {
          "display_name": "Python 3",
          "language": "python",
          "name": "python3"
        },
        "language_info": {
          "codemirror_mode": {
            "name": "ipython",
            "version": 3
          },
          "file_extension": ".py",
          "mimetype": "text/x-python",
          "name": "python",
          "nbconvert_exporter": "python",
          "pygments_lexer": "ipython3",
          "version": "3.8.5"
        }
      },
      "nbformat": 4,
      "nbformat_minor": 4
    }
    EOF
}

resource "google_colab_schedule" "schedule" {
  display_name = "tf-test-schedule%{random_suffix}"
  location = "%{location}"
  allow_queueing = true
  max_concurrent_run_count = 2
  cron = "TZ=America/Los_Angeles * * * * *"
  max_run_count = 5
  start_time = "%{start_time}"
  end_time = "%{end_time}"

  desired_state = "ACTIVE"

  create_notebook_execution_job_request {
    notebook_execution_job {
      display_name = "Notebook execution"
      gcs_notebook_source {
        uri = "gs://${google_storage_bucket_object.notebook.bucket}/${google_storage_bucket_object.notebook.name}"
        generation = google_storage_bucket_object.notebook.generation
      }

      notebook_runtime_template_resource_name = "projects/${google_colab_runtime_template.my_runtime_template.project}/locations/${google_colab_runtime_template.my_runtime_template.location}/notebookRuntimeTemplates/${google_colab_runtime_template.my_runtime_template.name}"
      gcs_output_uri = "gs://${google_storage_bucket.output_bucket.name}"
      service_account = "%{service_account}"
      }
  }

  depends_on = [
    google_colab_runtime_template.my_runtime_template,
    google_storage_bucket.output_bucket,
  ]
}
`, context)
}

func testAccColabSchedule_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_colab_runtime_template" "my_runtime_template" {
  name = "tf-test-runtime-template%{random_suffix}"
  display_name = "Runtime template"
  location = "us-central1"

  machine_spec {
    machine_type     = "e2-standard-4"
  }

  network_spec {
    enable_internet_access = true
  }
}

resource "google_storage_bucket" "output_bucket" {
  name          = "tf_test_my_bucket%{random_suffix}"
  location      = "US"
  force_destroy = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "notebook" {
  name   = "hello_world.ipynb"
  bucket = google_storage_bucket.output_bucket.name
  content = <<EOF
    {
      "cells": [
        {
          "cell_type": "code",
          "execution_count": null,
          "metadata": {},
          "outputs": [],
          "source": [
            "print(\"Hello, World!\")"
          ]
        }
      ],
      "metadata": {
        "kernelspec": {
          "display_name": "Python 3",
          "language": "python",
          "name": "python3"
        },
        "language_info": {
          "codemirror_mode": {
            "name": "ipython",
            "version": 3
          },
          "file_extension": ".py",
          "mimetype": "text/x-python",
          "name": "python",
          "nbconvert_exporter": "python",
          "pygments_lexer": "ipython3",
          "version": "3.8.5"
        }
      },
      "nbformat": 4,
      "nbformat_minor": 4
    }
    EOF
}

resource "google_colab_schedule" "schedule" {
  display_name = "tf-test-schedule-updated%{random_suffix}"
  location = "%{location}"
  allow_queueing = false
  max_concurrent_run_count = 1
  cron = "TZ=America/Los_Angeles 0 * * * *"
  max_run_count = 3
  start_time = "%{updated_start_time}"
  end_time = "%{updated_end_time}"

  create_notebook_execution_job_request {
    notebook_execution_job {
      display_name = "Notebook execution"
      gcs_notebook_source {
        uri = "gs://${google_storage_bucket_object.notebook.bucket}/${google_storage_bucket_object.notebook.name}"
        generation = google_storage_bucket_object.notebook.generation
      }

      notebook_runtime_template_resource_name = "projects/${google_colab_runtime_template.my_runtime_template.project}/locations/${google_colab_runtime_template.my_runtime_template.location}/notebookRuntimeTemplates/${google_colab_runtime_template.my_runtime_template.name}"
      gcs_output_uri = "gs://${google_storage_bucket.output_bucket.name}"
      service_account = "%{service_account}"
      }
  }

  depends_on = [
    google_colab_runtime_template.my_runtime_template,
    google_storage_bucket.output_bucket,
  ]
}
`, context)
}
