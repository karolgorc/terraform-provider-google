---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/compute_vpn_gateway.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  Get a VPN gateway within GCE.
---

# google_compute_vpn_gateway

Get a VPN gateway within GCE from its name.

## Example Usage

```tf
data "google_compute_vpn_gateway" "my-vpn-gateway" {
  name = "vpn-gateway-us-east1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the VPN gateway.


- - -

* `project` - (Optional) The project in which the resource belongs. If it
    is not provided, the provider project is used.

* `region` - (Optional) The region in which the resource belongs. If it
    is not provided, the project region is used.

## Attributes Reference

In addition to the arguments listed above, the following attributes are exported:

* `network` - The network of this VPN gateway.

* `description` - Description of this VPN gateway.

* `region` - Region of this VPN gateway.

* `self_link` - The URI of the resource.
