---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/looker/Instance.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Looker (Google Cloud core)"
description: |-
  A Google Cloud Looker instance.
---

# google_looker_instance

A Google Cloud Looker instance.


To get more information about Instance, see:

* [API documentation](https://cloud.google.com/looker/docs/reference/rest/v1/projects.locations.instances)
* How-to Guides
    * [Configure a Looker (Google Cloud core) instance](https://cloud.google.com/looker/docs/looker-core-instance-setup)
    * [Create a Looker (Google Cloud core) instance](https://cloud.google.com/looker/docs/looker-core-instance-create)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=looker_instance_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Looker Instance Basic


```hcl
resource "google_looker_instance" "looker-instance" {
  name              = "my-instance"
  platform_edition  = "LOOKER_CORE_STANDARD_ANNUAL"
  region            = "us-central1"
  oauth_config {
    client_id = "my-client-id"
    client_secret = "my-client-secret"
  }
  deletion_policy = "DEFAULT"
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=looker_instance_full&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Looker Instance Full


```hcl
resource "google_looker_instance" "looker-instance" {
  name               = "my-instance"
  platform_edition   = "LOOKER_CORE_STANDARD_ANNUAL"
  region             = "us-central1"
  public_ip_enabled  = true
  admin_settings {
    allowed_email_domains = ["google.com"]
  }
  maintenance_window {
    day_of_week = "THURSDAY"
    start_time {
      hours   = 22
      minutes = 0
      seconds = 0
      nanos   = 0
    }
  }
  deny_maintenance_period {    
    start_date {
      year = 2050
      month = 1
      day = 1
    }
    end_date {
      year = 2050
      month = 2
      day = 1
    }
    time {
      hours = 10
      minutes = 0
      seconds = 0
      nanos = 0
    }
  }
  oauth_config {
    client_id = "my-client-id"
    client_secret = "my-client-secret"
  }  
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=looker_instance_fips&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Looker Instance Fips


```hcl
resource "google_looker_instance" "looker-instance" {
  name               = "my-instance-fips"
  platform_edition   = "LOOKER_CORE_ENTERPRISE_ANNUAL"
  region             = "us-central1"
  public_ip_enabled  = true
  fips_enabled = true
  oauth_config {
    client_id = "my-client-id"
    client_secret = "my-client-secret"
  }  
}
```
## Example Usage - Looker Instance Enterprise Full


```hcl
resource "google_looker_instance" "looker-instance" {
  name               = "my-instance"
  platform_edition   = "LOOKER_CORE_ENTERPRISE_ANNUAL"
  region             = "us-central1"
  private_ip_enabled = true
  public_ip_enabled  = false
  reserved_range     = "${google_compute_global_address.looker_range.name}"
  consumer_network   = google_compute_network.looker_network.id
  admin_settings {
    allowed_email_domains = ["google.com"]
  }
  encryption_config {
    kms_key_name = "looker-kms-key"
  }
  maintenance_window {
    day_of_week = "THURSDAY"
    start_time {
      hours   = 22
      minutes = 0
      seconds = 0
      nanos   = 0
    }
  }
  deny_maintenance_period {
    start_date {
      year = 2050
      month = 1
      day = 1
    }
    end_date {
      year = 2050
      month = 2
      day = 1
    }
    time {
      hours = 10
      minutes = 0
      seconds = 0
      nanos = 0
    }
  }
  oauth_config {
    client_id = "my-client-id"
    client_secret = "my-client-secret"
  }
  depends_on   = [
    google_service_networking_connection.looker_vpc_connection
  ]
}

resource "google_service_networking_connection" "looker_vpc_connection" {
  network                 = google_compute_network.looker_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.looker_range.name]
}

resource "google_compute_global_address" "looker_range" {
  name          = "looker-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 20
  network       = google_compute_network.looker_network.id
}

data "google_project" "project" {}

resource "google_compute_network" "looker_network" {
  name = "looker-network"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "looker-kms-key"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-looker.iam.gserviceaccount.com"
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=looker_instance_custom_domain&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Looker Instance Custom Domain


```hcl
resource "google_looker_instance" "looker-instance" {
  name              = "my-instance"
  platform_edition  = "LOOKER_CORE_STANDARD_ANNUAL"
  region            = "us-central1"
  oauth_config {
    client_id = "my-client-id"
    client_secret = "my-client-secret"
  }
  // After your Looker (Google Cloud core) instance has been created, you can set up, view information about, or delete a custom domain for your instance. 
  // Therefore 2 terraform applies, one to create the instance, then another to set up the custom domain. 
  custom_domain {
    domain = "my-custom-domain.com"
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=looker_instance_psc&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Looker Instance Psc


```hcl
resource "google_looker_instance" "looker-instance" {
  name               = "my-instance"
  platform_edition   = "LOOKER_CORE_ENTERPRISE_ANNUAL"
  region             = "us-central1"
  private_ip_enabled = false
  public_ip_enabled  = false
  psc_enabled        = true
  oauth_config {
    client_id = "my-client-id"
    client_secret = "my-client-secret"
  }
  psc_config {
    allowed_vpcs = ["projects/test-project/global/networks/test"]
    # update only
    # service_attachments = [{local_fqdn: "www.local-fqdn.com" target_service_attachment_uri: "projects/my-project/regions/us-east1/serviceAttachments/sa"}]
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=looker_instance_force_delete&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Looker Instance Force Delete


```hcl
resource "google_looker_instance" "looker-instance" {
  name              = "my-instance"
  platform_edition  = "LOOKER_CORE_STANDARD_ANNUAL"
  region            = "us-central1"
  oauth_config {
    client_id = "my-client-id"
    client_secret = "my-client-secret"
  }
  deletion_policy = "FORCE"
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  The ID of the instance or a fully qualified identifier for the instance.

* `oauth_config` -
  (Required)
  Looker Instance OAuth login settings.
  Structure is [documented below](#nested_oauth_config).


* `admin_settings` -
  (Optional)
  Looker instance Admin settings.
  Structure is [documented below](#nested_admin_settings).

* `consumer_network` -
  (Optional)
  Network name in the consumer project in the format of: projects/{project}/global/networks/{network}
  Note that the consumer network may be in a different GCP project than the consumer
  project that is hosting the Looker Instance.

* `deny_maintenance_period` -
  (Optional)
  Maintenance denial period for this instance.
  You must allow at least 14 days of maintenance availability
  between any two deny maintenance periods.
  Structure is [documented below](#nested_deny_maintenance_period).

* `encryption_config` -
  (Optional)
  Looker instance encryption settings.
  Structure is [documented below](#nested_encryption_config).

* `fips_enabled` -
  (Optional)
  FIPS 140-2 Encryption enablement for Looker (Google Cloud Core).

* `maintenance_window` -
  (Optional)
  Maintenance window for an instance.
  Maintenance of your instance takes place once a month, and will require
  your instance to be restarted during updates, which will temporarily
  disrupt service.
  Structure is [documented below](#nested_maintenance_window).

* `platform_edition` -
  (Optional)
  Platform editions for a Looker instance. Each edition maps to a set of instance features, like its size. Must be one of these values:
  - LOOKER_CORE_TRIAL: trial instance (Currently Unavailable)
  - LOOKER_CORE_STANDARD: pay as you go standard instance (Currently Unavailable)
  - LOOKER_CORE_STANDARD_ANNUAL: subscription standard instance
  - LOOKER_CORE_ENTERPRISE_ANNUAL: subscription enterprise instance
  - LOOKER_CORE_EMBED_ANNUAL: subscription embed instance
  - LOOKER_CORE_NONPROD_STANDARD_ANNUAL: nonprod subscription standard instance
  - LOOKER_CORE_NONPROD_ENTERPRISE_ANNUAL: nonprod subscription enterprise instance
  - LOOKER_CORE_NONPROD_EMBED_ANNUAL: nonprod subscription embed instance
  - LOOKER_CORE_TRIAL_STANDARD: A standard trial edition of Looker (Google Cloud core) product.
  - LOOKER_CORE_TRIAL_ENTERPRISE: An enterprise trial edition of Looker (Google Cloud core) product.
  - LOOKER_CORE_TRIAL_EMBED: An embed trial edition of Looker (Google Cloud core) product.
  Default value is `LOOKER_CORE_TRIAL`.
  Possible values are: `LOOKER_CORE_TRIAL`, `LOOKER_CORE_STANDARD`, `LOOKER_CORE_STANDARD_ANNUAL`, `LOOKER_CORE_ENTERPRISE_ANNUAL`, `LOOKER_CORE_EMBED_ANNUAL`, `LOOKER_CORE_NONPROD_STANDARD_ANNUAL`, `LOOKER_CORE_NONPROD_ENTERPRISE_ANNUAL`, `LOOKER_CORE_NONPROD_EMBED_ANNUAL`, `LOOKER_CORE_TRIAL_STANDARD`, `LOOKER_CORE_TRIAL_ENTERPRISE`, `LOOKER_CORE_TRIAL_EMBED`.

* `private_ip_enabled` -
  (Optional)
  Whether private IP is enabled on the Looker instance.

* `psc_config` -
  (Optional)
  Information for Private Service Connect (PSC) setup for a Looker instance.
  Structure is [documented below](#nested_psc_config).

* `psc_enabled` -
  (Optional)
  Whether Public Service Connect (PSC) is enabled on the Looker instance

* `public_ip_enabled` -
  (Optional)
  Whether public IP is enabled on the Looker instance.

* `reserved_range` -
  (Optional)
  Name of a reserved IP address range within the consumer network, to be used for
  private service access connection. User may or may not specify this in a request.

* `user_metadata` -
  (Optional)
  Metadata about users for a Looker instance.
  These settings are only available when platform edition LOOKER_CORE_STANDARD is set.
  There are ten Standard and two Developer users included in the cost of the product.
  You can allocate additional Standard, Viewer, and Developer users for this instance.
  It is an optional step and can be modified later.
  With the Standard edition of Looker (Google Cloud core), you can provision up to 50
  total users, distributed across Viewer, Standard, and Developer.
  Structure is [documented below](#nested_user_metadata).

* `custom_domain` -
  (Optional)
  Custom domain settings for a Looker instance.
  Structure is [documented below](#nested_custom_domain).

* `region` -
  (Optional)
  The name of the Looker region of the instance.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.

* `deletion_policy` - (Optional) Policy to determine if the cluster should be deleted forcefully.
If setting deletion_policy = "FORCE", the Looker instance will be deleted regardless
of its nested resources. If set to "DEFAULT", Looker instances that still have
nested resources will return an error. Possible values: DEFAULT, FORCE



<a name="nested_oauth_config"></a>The `oauth_config` block supports:

* `client_id` -
  (Required)
  The client ID for the Oauth config.

* `client_secret` -
  (Required)
  The client secret for the Oauth config.

<a name="nested_admin_settings"></a>The `admin_settings` block supports:

* `allowed_email_domains` -
  (Optional)
  Email domain allowlist for the instance.
  Define the email domains to which your users can deliver Looker (Google Cloud core) content.
  Updating this list will restart the instance. Updating the allowed email domains from terraform
  means the value provided will be considered as the entire list and not an amendment to the
  existing list of allowed email domains.

<a name="nested_deny_maintenance_period"></a>The `deny_maintenance_period` block supports:

* `start_date` -
  (Required)
  Required. Start date of the deny maintenance period
  Structure is [documented below](#nested_deny_maintenance_period_start_date).

* `end_date` -
  (Required)
  Required. Start date of the deny maintenance period
  Structure is [documented below](#nested_deny_maintenance_period_end_date).

* `time` -
  (Required)
  Required. Start time of the window in UTC time.
  Structure is [documented below](#nested_deny_maintenance_period_time).


<a name="nested_deny_maintenance_period_start_date"></a>The `start_date` block supports:

* `year` -
  (Optional)
  Year of the date. Must be from 1 to 9999, or 0 to specify a date without
  a year.

* `month` -
  (Optional)
  Month of a year. Must be from 1 to 12, or 0 to specify a year without a
  month and day.

* `day` -
  (Optional)
  Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
  to specify a year by itself or a year and month where the day isn't significant.

<a name="nested_deny_maintenance_period_end_date"></a>The `end_date` block supports:

* `year` -
  (Optional)
  Year of the date. Must be from 1 to 9999, or 0 to specify a date without
  a year.

* `month` -
  (Optional)
  Month of a year. Must be from 1 to 12, or 0 to specify a year without a
  month and day.

* `day` -
  (Optional)
  Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
  to specify a year by itself or a year and month where the day isn't significant.

<a name="nested_deny_maintenance_period_time"></a>The `time` block supports:

* `hours` -
  (Optional)
  Hours of day in 24 hour format. Should be from 0 to 23.

* `minutes` -
  (Optional)
  Minutes of hour of day. Must be from 0 to 59.

* `seconds` -
  (Optional)
  Seconds of minutes of the time. Must normally be from 0 to 59.

* `nanos` -
  (Optional)
  Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.

<a name="nested_encryption_config"></a>The `encryption_config` block supports:

* `kms_key_name` -
  (Optional)
  Name of the customer managed encryption key (CMEK) in KMS.

* `kms_key_state` -
  (Output)
  Status of the customer managed encryption key (CMEK) in KMS.

* `kms_key_name_version` -
  (Output)
  Full name and version of the CMEK key currently in use to encrypt Looker data.

<a name="nested_maintenance_window"></a>The `maintenance_window` block supports:

* `day_of_week` -
  (Required)
  Required. Day of the week for this MaintenanceWindow (in UTC).
  - MONDAY: Monday
  - TUESDAY: Tuesday
  - WEDNESDAY: Wednesday
  - THURSDAY: Thursday
  - FRIDAY: Friday
  - SATURDAY: Saturday
  - SUNDAY: Sunday
  Possible values are: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`.

* `start_time` -
  (Required)
  Required. Start time of the window in UTC time.
  Structure is [documented below](#nested_maintenance_window_start_time).


<a name="nested_maintenance_window_start_time"></a>The `start_time` block supports:

* `hours` -
  (Optional)
  Hours of day in 24 hour format. Should be from 0 to 23.

* `minutes` -
  (Optional)
  Minutes of hour of day. Must be from 0 to 59.

* `seconds` -
  (Optional)
  Seconds of minutes of the time. Must normally be from 0 to 59.

* `nanos` -
  (Optional)
  Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.

<a name="nested_psc_config"></a>The `psc_config` block supports:

* `allowed_vpcs` -
  (Optional)
  List of VPCs that are allowed ingress into the Looker instance.

* `looker_service_attachment_uri` -
  (Output)
  URI of the Looker service attachment.

* `service_attachments` -
  (Optional)
  List of egress service attachment configurations.
  Structure is [documented below](#nested_psc_config_service_attachments).


<a name="nested_psc_config_service_attachments"></a>The `service_attachments` block supports:

* `connection_status` -
  (Output)
  Status of the service attachment connection.

* `local_fqdn` -
  (Optional)
  Fully qualified domain name that will be used in the private DNS record created for the service attachment.

* `target_service_attachment_uri` -
  (Optional)
  URI of the service attachment to connect to.

<a name="nested_user_metadata"></a>The `user_metadata` block supports:

* `additional_viewer_user_count` -
  (Optional)
  Number of additional Viewer Users to allocate to the Looker Instance.

* `additional_standard_user_count` -
  (Optional)
  Number of additional Standard Users to allocate to the Looker Instance.

* `additional_developer_user_count` -
  (Optional)
  Number of additional Developer Users to allocate to the Looker Instance.

<a name="nested_custom_domain"></a>The `custom_domain` block supports:

* `domain` -
  (Optional)
  Domain name

* `state` -
  (Output)
  Status of the custom domain.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{region}}/instances/{{name}}`

* `create_time` -
  The time the instance was created in RFC3339 UTC "Zulu" format,
  accurate to nanoseconds.

* `egress_public_ip` -
  Public Egress IP (IPv4).

* `ingress_private_ip` -
  Private Ingress IP (IPv4).

* `ingress_public_ip` -
  Public Ingress IP (IPv4).

* `looker_version` -
  The Looker version that the instance is using.

* `looker_uri` -
  Looker instance URI which can be used to access the Looker Instance UI.

* `update_time` -
  The time the instance was updated in RFC3339 UTC "Zulu" format,
  accurate to nanoseconds.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 90 minutes.
- `update` - Default is 90 minutes.
- `delete` - Default is 90 minutes.

## Import


Instance can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{region}}/instances/{{name}}`
* `{{project}}/{{region}}/{{name}}`
* `{{region}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Instance using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{region}}/instances/{{name}}"
  to = google_looker_instance.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Instance can be imported using one of the formats above. For example:

```
$ terraform import google_looker_instance.default projects/{{project}}/locations/{{region}}/instances/{{name}}
$ terraform import google_looker_instance.default {{project}}/{{region}}/{{name}}
$ terraform import google_looker_instance.default {{region}}/{{name}}
$ terraform import google_looker_instance.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
