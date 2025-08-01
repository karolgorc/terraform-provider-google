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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/compute/resource_compute_region_disk_test.go.tmpl
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package compute_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"

	"google.golang.org/api/compute/v1"
)

func TestAccComputeRegionDisk_basic(t *testing.T) {
	t.Parallel()

	diskName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	var disk compute.Disk

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_basic(diskName, "self_link"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
				),
			},
			{
				ResourceName:            "google_compute_region_disk.regiondisk",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
			{
				Config: testAccComputeRegionDisk_basic(diskName, "name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
				),
			},
			{
				ResourceName:            "google_compute_region_disk.regiondisk",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func TestAccComputeRegionDisk_hyperdisk(t *testing.T) {
	t.Parallel()

	diskName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	var disk compute.Disk

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_hyperdisk(diskName, "self_link"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
				),
			},
			{
				ResourceName:            "google_compute_region_disk.regiondisk",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
			{
				Config: testAccComputeRegionDisk_hyperdiskUpdated(diskName, "name"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Check that the update is done in-place
						plancheck.ExpectResourceAction("google_compute_region_disk.regiondisk", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_compute_region_disk.regiondisk", "access_mode", "READ_WRITE_SINGLE"),
					resource.TestCheckResourceAttr("google_compute_region_disk.regiondisk", "provisioned_iops", "20000"),
					resource.TestCheckResourceAttr("google_compute_region_disk.regiondisk", "provisioned_throughput", "250"),
					testAccCheckComputeRegionDiskExists(t, "google_compute_region_disk.regiondisk", &disk),
				),
			},
			{
				ResourceName:            "google_compute_region_disk.regiondisk",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func TestAccComputeRegionDisk_basicUpdate(t *testing.T) {
	t.Parallel()

	diskName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	var disk compute.Disk

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_basic(diskName, "self_link"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
				),
			},
			{
				ResourceName:            "google_compute_region_disk.regiondisk",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
			{
				Config: testAccComputeRegionDisk_basicUpdated(diskName, "self_link"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
					resource.TestCheckResourceAttr("google_compute_region_disk.regiondisk", "size", "100"),
					testAccCheckComputeRegionDiskHasLabel(&disk, "my-label", "my-updated-label-value"),
					testAccCheckComputeRegionDiskHasLabel(&disk, "a-new-label", "a-new-label-value"),
					testAccCheckComputeRegionDiskHasLabelFingerprint(&disk, "google_compute_region_disk.regiondisk"),
				),
			},
			{
				ResourceName:            "google_compute_region_disk.regiondisk",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func TestAccComputeRegionDisk_encryption(t *testing.T) {
	t.Parallel()

	diskName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	var disk compute.Disk

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_encryption(diskName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
					testAccCheckRegionDiskEncryptionKey(
						"google_compute_region_disk.regiondisk", &disk),
				),
			},
		},
	})
}

func TestAccComputeRegionDisk_deleteDetach(t *testing.T) {
	t.Parallel()

	diskName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	regionDiskName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	regionDiskName2 := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	var disk compute.Disk

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_deleteDetach(instanceName, diskName, regionDiskName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
				),
			},
			// this needs to be an additional step so we refresh and see the instance
			// listed as attached to the disk; the instance is created after the
			// disk. and the disk's properties aren't refreshed unless there's
			// another step
			{
				Config: testAccComputeRegionDisk_deleteDetach(instanceName, diskName, regionDiskName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
					testAccCheckComputeRegionDiskInstances(
						"google_compute_region_disk.regiondisk", &disk),
				),
			},
			// Change the disk name to destroy it, which detaches it from the instance
			{
				Config: testAccComputeRegionDisk_deleteDetach(instanceName, diskName, regionDiskName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
				),
			},
			// Add the extra step like before
			{
				Config: testAccComputeRegionDisk_deleteDetach(instanceName, diskName, regionDiskName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
					testAccCheckComputeRegionDiskInstances(
						"google_compute_region_disk.regiondisk", &disk),
				),
			},
		},
	})
}

func TestAccComputeRegionDisk_cloneDisk(t *testing.T) {
	t.Parallel()

	diskName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	var disk compute.Disk

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_diskClone(diskName, "self_link"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk-clone", &disk),
				),
			},
			{
				ResourceName:      "google_compute_region_disk.regiondisk-clone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionDisk_featuresUpdated(t *testing.T) {
	t.Parallel()

	diskName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	var disk compute.Disk

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_features(diskName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.regiondisk", &disk),
				),
			},
			{
				ResourceName:      "google_compute_region_disk.regiondisk",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionDisk_featuresUpdated(diskName),
			},
			{
				ResourceName:      "google_compute_region_disk.regiondisk",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionDisk_createSnapshotBeforeDestroy(t *testing.T) {
	acctest.SkipIfVcr(t) // Disk cleanup test check
	t.Parallel()

	var disk1 compute.Disk
	var disk2 compute.Disk
	var disk3 compute.Disk
	context := map[string]interface{}{
		"disk_name1":        fmt.Sprintf("tf-test-disk-%s", acctest.RandString(t, 10)),
		"disk_name2":        fmt.Sprintf("test-%s", acctest.RandString(t, 44)), //this is over the snapshot character creation limit of 48
		"disk_name3":        fmt.Sprintf("tf-test-disk-%s", acctest.RandString(t, 10)),
		"snapshot_prefix":   fmt.Sprintf("tf-test-snapshot-%s", acctest.RandString(t, 10)),
		"kms_key_self_link": acctest.BootstrapKMSKey(t).CryptoKey.Name,
		"raw_key":           "SGVsbG8gZnJvbSBHb29nbGUgQ2xvdWQgUGxhdGZvcm0=",
		"rsa_encrypted_key": "ieCx/NcW06PcT7Ep1X6LUTc/hLvUDYyzSZPPVCVPTVEohpeHASqC8uw5TzyO9U+Fka9JFHz0mBibXUInrC/jEk014kCK/NPjYgEMOyssZ4ZINPKxlUh2zn1bV+MCaTICrdmuSBTWlUUiFoDD6PYznLwh8ZNdaheCeZ8ewEXgFQ8V+sDroLaN3Xs3MDTXQEMMoNUXMCZEIpg9Vtp9x2oeQ5lAbtt7bYAAHf5l+gJWw3sUfs0/Glw5fpdjT8Uggrr+RMZezGrltJEF293rvTIjWOEB3z5OHyHwQkvdrPDFcTqsLfh+8Hr8g+mf+7zVPEC8nEbqpdl3GPv3A7AwpFp7MA==",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionDiskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDisk_createSnapshotBeforeDestroy_init(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.raw-encrypted-name", &disk1),
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.rsa-encrypted-prefix", &disk2),
					testAccCheckComputeRegionDiskExists(
						t, "google_compute_region_disk.kms-encrypted-name", &disk3),
				),
			},
			{
				Config:  testAccComputeRegionDisk_createSnapshotBeforeDestroy_init(context),
				Destroy: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeDisk_removeBackupSnapshot(t, context["disk_name1"].(string)),
					testAccCheckComputeDisk_removeBackupSnapshot(t, context["snapshot_prefix"].(string)),
					testAccCheckComputeDisk_removeBackupSnapshot(t, context["disk_name3"].(string)),
				),
			},
		},
	})
}

func testAccCheckComputeRegionDiskExists(t *testing.T, n string, disk *compute.Disk) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		p := envvar.GetTestProjectFromEnv()
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.Attributes["name"] == "" {
			return fmt.Errorf("No ID is set")
		}

		config := acctest.GoogleProviderConfig(t)

		found, err := config.NewComputeClient(config.UserAgent).RegionDisks.Get(
			p, rs.Primary.Attributes["region"], rs.Primary.Attributes["name"]).Do()
		if err != nil {
			return err
		}

		if found.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("RegionDisk not found")
		}

		*disk = *found

		return nil
	}
}

func testAccCheckComputeRegionDiskHasLabel(disk *compute.Disk, key, value string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		val, ok := disk.Labels[key]
		if !ok {
			return fmt.Errorf("Label with key %s not found", key)
		}

		if val != value {
			return fmt.Errorf("Label value did not match for key %s: expected %s but found %s", key, value, val)
		}
		return nil
	}
}

func testAccCheckComputeRegionDiskHasLabelFingerprint(disk *compute.Disk, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		state := s.RootModule().Resources[resourceName]
		if state == nil {
			return fmt.Errorf("Unable to find resource named %s", resourceName)
		}

		labelFingerprint := state.Primary.Attributes["label_fingerprint"]
		if labelFingerprint != disk.LabelFingerprint {
			return fmt.Errorf("Label fingerprints do not match: api returned %s but state has %s",
				disk.LabelFingerprint, labelFingerprint)
		}

		return nil
	}
}

func testAccCheckRegionDiskEncryptionKey(n string, disk *compute.Disk) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		attr := rs.Primary.Attributes["disk_encryption_key.0.sha256"]
		if disk.DiskEncryptionKey == nil {
			return fmt.Errorf("RegionDisk %s has mismatched encryption key.\nTF State: %+v\nGCP State: <empty>", n, attr)
		} else if attr != disk.DiskEncryptionKey.Sha256 {
			return fmt.Errorf("RegionDisk %s has mismatched encryption key.\nTF State: %+v.\nGCP State: %+v",
				n, attr, disk.DiskEncryptionKey.Sha256)
		}
		return nil
	}
}

func testAccCheckComputeRegionDiskInstances(n string, disk *compute.Disk) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		attr := rs.Primary.Attributes["users.#"]
		if strconv.Itoa(len(disk.Users)) != attr {
			return fmt.Errorf("RegionDisk %s has mismatched users.\nTF State: %+v\nGCP State: %+v", n, rs.Primary.Attributes["users"], disk.Users)
		}

		for pos, user := range disk.Users {
			if tpgresource.ConvertSelfLinkToV1(rs.Primary.Attributes["users."+strconv.Itoa(pos)]) != tpgresource.ConvertSelfLinkToV1(user) {
				return fmt.Errorf("RegionDisk %s has mismatched users.\nTF State: %+v.\nGCP State: %+v",
					n, rs.Primary.Attributes["users"], disk.Users)
			}
		}
		return nil
	}
}

func testAccComputeRegionDisk_hyperdisk(diskName, refSelector string) string {
	return fmt.Sprintf(`
resource "google_compute_disk" "disk" {
	name  = "%s"
	image = "debian-cloud/debian-11"
	size  = 50
	type  = "pd-ssd"
	zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
	name        = "%s"
	source_disk = google_compute_disk.disk.name
	zone        = "us-central1-a"
}

resource "google_compute_region_disk" "regiondisk" {
	name     = "%s"
	snapshot = google_compute_snapshot.snapdisk.%s
	type     = "hyperdisk-balanced-high-availability"
	size     = 50
	replica_zones = ["us-central1-a", "us-central1-f"]

	access_mode            = "READ_WRITE_MANY"
	provisioned_iops       = 10000
	provisioned_throughput = 190
}
`, diskName, diskName, diskName, refSelector)
}

func testAccComputeRegionDisk_hyperdiskUpdated(diskName, refSelector string) string {
	return fmt.Sprintf(`
resource "google_compute_disk" "disk" {
	name  = "%s"
	image = "debian-cloud/debian-11"
	size  = 50
	type  = "pd-ssd"
	zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
	name        = "%s"
	source_disk = google_compute_disk.disk.name
	zone        = "us-central1-a"
}

resource "google_compute_region_disk" "regiondisk" {
	name     = "%s"
	snapshot = google_compute_snapshot.snapdisk.%s
	type     = "hyperdisk-balanced-high-availability"
	region   = "us-central1"

	replica_zones = ["us-central1-a", "us-central1-f"]

	size = 100
	access_mode            = "READ_WRITE_SINGLE"
	provisioned_iops       = 20000
	provisioned_throughput = 250
}
`, diskName, diskName, diskName, refSelector)
}

func testAccComputeRegionDisk_basic(diskName, refSelector string) string {
	return fmt.Sprintf(`
resource "google_compute_disk" "disk" {
  name  = "%s"
  image = "debian-cloud/debian-11"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name        = "%s"
  source_disk = google_compute_disk.disk.name
  zone        = "us-central1-a"
}

resource "google_compute_region_disk" "regiondisk" {
  name     = "%s"
  snapshot = google_compute_snapshot.snapdisk.%s
  type     = "pd-ssd"
  replica_zones = ["us-central1-a", "us-central1-f"]
}
`, diskName, diskName, diskName, refSelector)
}

func testAccComputeRegionDisk_basicUpdated(diskName, refSelector string) string {
	return fmt.Sprintf(`
resource "google_compute_disk" "disk" {
  name  = "%s"
  image = "debian-cloud/debian-11"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name        = "%s"
  source_disk = google_compute_disk.disk.name
  zone        = "us-central1-a"
}

resource "google_compute_region_disk" "regiondisk" {
  name     = "%s"
  snapshot = google_compute_snapshot.snapdisk.%s
  type     = "pd-ssd"
  region   = "us-central1"

  replica_zones = ["us-central1-a", "us-central1-f"]

  size = 100
  labels = {
    my-label    = "my-updated-label-value"
    a-new-label = "a-new-label-value"
  }
}
`, diskName, diskName, diskName, refSelector)
}

func testAccComputeRegionDisk_encryption(diskName string) string {
	return fmt.Sprintf(`
resource "google_compute_disk" "disk" {
  name  = "%s"
  image = "debian-cloud/debian-11"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name = "%s"
  zone = "us-central1-a"

  source_disk = google_compute_disk.disk.name
}

resource "google_compute_region_disk" "regiondisk" {
  name     = "%s"
  snapshot = google_compute_snapshot.snapdisk.self_link
  type     = "pd-ssd"

  replica_zones = ["us-central1-a", "us-central1-f"]

  disk_encryption_key {
    raw_key = "SGVsbG8gZnJvbSBHb29nbGUgQ2xvdWQgUGxhdGZvcm0="
  }
}
`, diskName, diskName, diskName)
}

func testAccComputeRegionDisk_deleteDetach(instanceName, diskName, regionDiskName string) string {
	return fmt.Sprintf(`
resource "google_compute_disk" "disk" {
  name  = "%s"
  image = "debian-cloud/debian-11"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name        = "%s"
  source_disk = google_compute_disk.disk.name
  zone        = "us-central1-a"
}

resource "google_compute_region_disk" "regiondisk" {
  name     = "%s"
  snapshot = google_compute_snapshot.snapdisk.self_link
  type     = "pd-ssd"

  replica_zones = ["us-central1-a", "us-central1-f"]
}

resource "google_compute_instance" "inst" {
  name         = "%s"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  attached_disk {
    source = google_compute_region_disk.regiondisk.self_link
  }

  network_interface {
    network = "default"
  }
}
`, diskName, diskName, regionDiskName, instanceName)
}

func testAccComputeRegionDisk_diskClone(diskName, refSelector string) string {
	return fmt.Sprintf(`
	  resource "google_compute_region_disk" "regiondisk" {
		name                      = "%s"
		snapshot                  = google_compute_snapshot.snapdisk.id
		type                      = "pd-ssd"
		region                    = "us-central1"
		physical_block_size_bytes = 4096
	  
		replica_zones = ["us-central1-a", "us-central1-f"]  
	  }
	  
	  resource "google_compute_disk" "disk" {
		name  = "%s"
		image = "debian-11-bullseye-v20220719"
		size  = 50
		type  = "pd-ssd"
		zone  = "us-central1-a"
	  }
	  
	  resource "google_compute_snapshot" "snapdisk" {
		name        = "%s"
		source_disk = google_compute_disk.disk.name
		zone        = "us-central1-a"
	  }

	  resource "google_compute_region_disk" "regiondisk-clone" {
		name                      = "%s"
		source_disk = google_compute_region_disk.regiondisk.%s
		type                      = "pd-ssd"
		region                    = "us-central1"
		physical_block_size_bytes = 4096
	  
		replica_zones = ["us-central1-a", "us-central1-f"]
	  }
	`, diskName, diskName, diskName, diskName+"-clone", refSelector)
}

func testAccComputeRegionDisk_features(diskName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_disk" "regiondisk" {
  name   = "%s"
  type   = "pd-ssd"
  size   = 50
  region = "us-central1"

  guest_os_features {
    type = "SECURE_BOOT"
  }

  replica_zones = ["us-central1-a", "us-central1-f"]
}
`, diskName)
}

func testAccComputeRegionDisk_featuresUpdated(diskName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_disk" "regiondisk" {
  name   = "%s"
  type   = "pd-ssd"
  size   = 50
  region = "us-central1"

  guest_os_features {
    type = "SECURE_BOOT"
  }

  guest_os_features {
    type = "MULTI_IP_SUBNET"
  }

  replica_zones = ["us-central1-a", "us-central1-f"]
}
`, diskName)
}

func testAccComputeRegionDisk_createSnapshotBeforeDestroy_init(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_disk" "raw-encrypted-name" {
  name = "%{disk_name1}"
  type = "pd-ssd"
  size = 10
  region = "us-central1"
  replica_zones = ["us-central1-a", "us-central1-f"]

  disk_encryption_key {
	raw_key = "%{raw_key}"
  }

  create_snapshot_before_destroy = true
}

resource "google_compute_region_disk" "rsa-encrypted-prefix" {
  name = "%{disk_name2}"
  type = "pd-ssd"
  size = 10
  region = "us-central1"
  replica_zones = ["us-central1-a", "us-central1-f"]

  disk_encryption_key {
	rsa_encrypted_key = "%{rsa_encrypted_key}"
  }

  create_snapshot_before_destroy = true
  create_snapshot_before_destroy_prefix = "%{snapshot_prefix}"
}

resource "google_compute_region_disk" "kms-encrypted-name" {
  name = "%{disk_name3}"
  type = "pd-ssd"
  size = 10
  region = "us-central1"
  replica_zones = ["us-central1-a", "us-central1-f"]

  disk_encryption_key {
	kms_key_name = "%{kms_key_self_link}"
  }

  create_snapshot_before_destroy = true
}

`, context)
}
