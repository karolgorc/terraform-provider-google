---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/sourcerepo/Repository.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/datasource_iam.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Source Repositories"
description: |-
  A datasource to retrieve the IAM policy state for Cloud Source Repositories Repository
---


# google_sourcerepo_repository_iam_policy

Retrieves the current IAM policy data for repository


## Example Usage


```hcl
data "google_sourcerepo_repository_iam_policy" "policy" {
  project = google_sourcerepo_repository.my-repo.project
  repository = google_sourcerepo_repository.my-repo.name
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) Used to find the parent resource to bind the IAM policy to

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.

## Attributes Reference

The attributes are exported:

* `etag` - (Computed) The etag of the IAM policy.

* `policy_data` - (Required only by `google_sourcerepo_repository_iam_policy`) The policy data generated by
  a `google_iam_policy` data source.
