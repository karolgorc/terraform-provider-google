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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/netapp/resource_netapp_storage_pool_test.go.tmpl
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package netapp_test

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccNetappStoragePool_storagePoolCreateExample_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "gcnv-network-config-3", acctest.ServiceNetworkWithParentService("netapp.servicenetworking.goog")),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetappStoragePool_storagePoolCreateExample_full(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
			{
				Config: testAccNetappStoragePool_storagePoolCreateExample_update(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetappStoragePool_storagePoolCreateExample_full(context map[string]interface{}) string {
	return acctest.Nprintf(`

data "google_compute_network" "default" {
    name = "%{network_name}"
}

resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "us-central1"
  service_level = "PREMIUM"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
  active_directory      = ""
  description           = "this is a test description"
  kms_config            = ""
  labels                = {
    key= "test"
    value= "pool"
  }
  ldap_enabled          = false

}
`, context)
}

func testAccNetappStoragePool_storagePoolCreateExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`

data "google_compute_network" "default" {
    name = "%{network_name}"
}

resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "us-central1"
  service_level = "PREMIUM"
  capacity_gib = "4096"
  network = data.google_compute_network.default.id
  active_directory      = ""
  description           = "this is test"
  kms_config            = ""
  labels                = {
    key= "test"
    value= "pool"
  }
  ldap_enabled          = false

}
`, context)
}

func TestAccNetappStoragePool_autoTieredStoragePoolCreateExample_update(t *testing.T) {
	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "gcnv-network-config-3", acctest.ServiceNetworkWithParentService("netapp.servicenetworking.goog")),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetappStoragePoolDestroyProducer(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNetappStoragePool_autoTieredStoragePoolCreateExample_full(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
			{
				Config: testAccNetappStoragePool_autoTieredStoragePoolCreateExample_update(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetappStoragePool_autoTieredStoragePoolCreateExample_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_network" "default" {
    name = "%{network_name}"
}

resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "us-east4"
  service_level = "PREMIUM"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
  active_directory      = ""
  description           = "this is a test description"
  kms_config            = ""
  labels                = {
    key= "test"
    value= "pool"
  }
  ldap_enabled          = false
  allow_auto_tiering    = false
}
`, context)
}

func testAccNetappStoragePool_autoTieredStoragePoolCreateExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_network" "default" {
    name = "%{network_name}"
}

resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "us-east4"
  service_level = "PREMIUM"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
  active_directory      = ""
  description           = "this is a test description"
  kms_config            = ""
  labels                = {
    key= "test"
    value= "pool"
  }
  ldap_enabled          = false
  allow_auto_tiering = true
}
`, context)
}

func TestAccNetappStoragePool_FlexRegionalStoragePoolCreateExample_update(t *testing.T) {
	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "gcnv-network-config-3", acctest.ServiceNetworkWithParentService("netapp.servicenetworking.goog")),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetappStoragePoolDestroyProducer(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNetappStoragePool_FlexRegionalStoragePoolCreateExample_full(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
			{
				Config: testAccNetappStoragePool_FlexRegionalStoragePoolCreateExample_switchZone(context),
				Check:  testAccNetappStoragePool_FlexRegionalStoragePoolCreateExample_sleep_5_mins(),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
			{
				Config: testAccNetappStoragePool_FlexRegionalStoragePoolCreateExample_switchBackZone(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetappStoragePool_FlexRegionalStoragePoolCreateExample_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "us-east1"
  service_level = "FLEX"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
  zone = "us-east1-c"
  replica_zone = "us-east1-b"
}

resource "time_sleep" "wait_5_minutes" {
    depends_on = [google_netapp_storage_pool.test_pool]
    destroy_duration = "5m"
}

data "google_compute_network" "default" {
    name = "%{network_name}"
}
`, context)
}

func testAccNetappStoragePool_FlexRegionalStoragePoolCreateExample_switchZone(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "us-east1"
  service_level = "FLEX"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
  zone = "us-east1-b"
  replica_zone = "us-east1-c"
}

resource "time_sleep" "wait_5_minutes" {
    depends_on = [google_netapp_storage_pool.test_pool]
    destroy_duration = "5m"
}

data "google_compute_network" "default" {
    name = "%{network_name}"
}
`, context)
}

func testAccNetappStoragePool_FlexRegionalStoragePoolCreateExample_sleep_5_mins() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// wait 5 minutes before executing the switchback due to api zone switch issues
		time.Sleep(5 * time.Minute)
		return nil
	}
}

func testAccNetappStoragePool_FlexRegionalStoragePoolCreateExample_switchBackZone(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "us-east1"
  service_level = "FLEX"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
  zone = "us-east1-c"
  replica_zone = "us-east1-b"
}

resource "time_sleep" "wait_5_minutes" {
    depends_on = [google_netapp_storage_pool.test_pool]
    destroy_duration = "5m"
}

data "google_compute_network" "default" {
    name = "%{network_name}"
}
`, context)
}

func TestAccNetappStoragePool_FlexRegionalStoragePoolNoZone(t *testing.T) {
	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "gcnv-network-config-3", acctest.ServiceNetworkWithParentService("netapp.servicenetworking.goog")),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetappStoragePoolDestroyProducer(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNetappStoragePool_FlexRegionalStoragePoolNoZone(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetappStoragePool_FlexRegionalStoragePoolNoZone(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "europe-west3-a"
  service_level = "FLEX"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
}

resource "time_sleep" "wait_5_minutes" {
    depends_on = [google_netapp_storage_pool.test_pool]
    destroy_duration = "5m"
}

data "google_compute_network" "default" {
    name = "%{network_name}"
}
`, context)
}

func TestAccNetappStoragePool_customPerformanceStoragePoolCreateExample_update(t *testing.T) {
	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "gcnv-network-config-3", acctest.ServiceNetworkWithParentService("netapp.servicenetworking.goog")),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetappStoragePoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetappStoragePool_customPerformanceStoragePoolCreateExample_full(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
			{
				Config: testAccNetappStoragePool_customPerformanceStoragePoolCreateExample_update(context),
			},
			{
				ResourceName:            "google_netapp_storage_pool.test_pool",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetappStoragePool_customPerformanceStoragePoolCreateExample_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "us-east4-a"
  service_level = "FLEX"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
  description = "this is a test description"
  custom_performance_enabled = true
  total_throughput_mibps = "64"
  total_iops = "1024"
}

data "google_compute_network" "default" {
    name = "%{network_name}"
}
`, context)
}

func testAccNetappStoragePool_customPerformanceStoragePoolCreateExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_netapp_storage_pool" "test_pool" {
  name = "tf-test-pool%{random_suffix}"
  location = "us-east4-a"
  service_level = "FLEX"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
  description = "this is updated test description"
  custom_performance_enabled = true
  total_throughput_mibps = "200"
  total_iops = "3200"
}

data "google_compute_network" "default" {
    name = "%{network_name}"
}
`, context)
}
