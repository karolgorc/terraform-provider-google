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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/storagebatchoperations/resource_storage_batch_operations_job_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package storagebatchoperations_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccStorageBatchOperationsJobs_createJobWithPrefix(t *testing.T) {
	t.Parallel()
	bucketName := acctest.TestBucketName(t)
	jobID := fmt.Sprintf("tf-test-job-%d", acctest.RandInt(t))
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBatchOperationsCreateJobWithPrefix(bucketName, jobID),
			},
			{
				ResourceName:            "google_storage_batch_operations_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"job_id", "delete_protection"},
			},
		},
	})
}

func TestAccStorageBatchOperationsJobs_jobWithPrefixDeleteObjectAllVersions(t *testing.T) {
	t.Parallel()
	bucketName := acctest.TestBucketName(t)
	jobID := fmt.Sprintf("tf-test-job-%d", acctest.RandInt(t))
	liveObjectJobID := fmt.Sprintf("tf-test-job-%d", acctest.RandInt(t))
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBatchOperationsJobWithPrefixDeleteObjectAllVersions(bucketName, jobID),
			},
			{
				ResourceName:            "google_storage_batch_operations_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"job_id", "delete_protection"},
			},
			{
				Config: testAccStorageBatchOperationsJobWithPrefixDeleteLiveObject(bucketName, liveObjectJobID),
			},
			{
				ResourceName:            "google_storage_batch_operations_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"job_id", "delete_protection"},
			},
		},
	})
}

func TestAccStorageBatchOperationsJobs_jobWithPrefixObjectHold(t *testing.T) {
	t.Parallel()
	bucketName := acctest.TestBucketName(t)
	jobID := fmt.Sprintf("tf-test-job-%d", acctest.RandInt(t))
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBatchOperationsJobWithPrefixObjectHold(bucketName, jobID),
			},
			{
				ResourceName:            "google_storage_batch_operations_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"job_id", "delete_protection"},
			},
		},
	})
}

func TestAccStorageBatchOperationsJobs_jobWithPrefixObjectTemporaryHold(t *testing.T) {
	t.Parallel()
	bucketName := acctest.TestBucketName(t)
	jobID := fmt.Sprintf("tf-test-job-%d", acctest.RandInt(t))
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBatchOperationsJobWithPrefixObjectTemporaryHold(bucketName, jobID),
			},
			{
				ResourceName:            "google_storage_batch_operations_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"job_id", "delete_protection"},
			},
		},
	})
}

func TestAccStorageBatchOperationsJobs_createJobWithManifest(t *testing.T) {
	t.Parallel()
	bucketName := acctest.TestBucketName(t)
	jobID := fmt.Sprintf("tf-test-job-%d", acctest.RandInt(t))
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBatchOperationsCreateJobWithManifest(bucketName, jobID),
			},
			{
				ResourceName:            "google_storage_batch_operations_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"job_id", "delete_protection"},
			},
		},
	})
}

func TestAccStorageBatchOperationsJobs_batchOperationJobKmsKey(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"kms_key":     acctest.BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"job_id":      fmt.Sprintf("tf-test-job-%d", acctest.RandInt(t)),
		"object_name": fmt.Sprintf("tf-test-object-%d", acctest.RandInt(t)),
		"bucket_name": acctest.TestBucketName(t),
	}
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageBatchOperationsJobs_storageBatchOerationsJobKmsKey(context),
			},
			{
				ResourceName:            "google_storage_batch_operations_job.job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"job_id", "delete_protection"},
			},
		},
	})
}

func testAccStorageBatchOperationsCreateJobWithPrefix(bucketName, jobID string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "%s"
  location = "us-central1"
  uniform_bucket_level_access = true
  force_destroy = true
}
resource "google_storage_batch_operations_job" "job" {
	job_id   = "%s"
	bucket_list {
		buckets  {
			bucket = google_storage_bucket.bucket.name
			prefix_list {
				included_object_prefixes = [
					"bkt"
				]
			}
		}
	}

	put_metadata {
		custom_metadata = {
			"key"="value"
			"key1"="value1"
		}
	}
	delete_protection = false
}
`, bucketName, jobID)
}

func testAccStorageBatchOperationsJobWithPrefixDeleteObjectAllVersions(bucketName, jobID string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "%s"
  location = "us-central1"
  uniform_bucket_level_access = true
  force_destroy = true
}
resource "google_storage_batch_operations_job" "job" {
	job_id   = "%s"
	bucket_list {
		buckets  {
			bucket = google_storage_bucket.bucket.name
			prefix_list {
				included_object_prefixes = [
					"bkt"
				]
			}
		}
	}
	delete_object {
		permanent_object_deletion_enabled = true
	}
	delete_protection = false
}
`, bucketName, jobID)
}

func testAccStorageBatchOperationsJobWithPrefixDeleteLiveObject(bucketName, jobID string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "%s"
  location = "us-central1"
  uniform_bucket_level_access = true
  force_destroy = true
}
resource "google_storage_batch_operations_job" "job" {
	job_id   = "%s"
	bucket_list {
		buckets  {
			bucket = google_storage_bucket.bucket.name
			prefix_list {
				included_object_prefixes = [
					"objprefix"
				]
			}
		}
	}
	delete_object {
		permanent_object_deletion_enabled = false
	}
	delete_protection = false
}
`, bucketName, jobID)
}

func testAccStorageBatchOperationsJobWithPrefixObjectHold(bucketName, jobID string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "%s"
  location = "us-central1"
  uniform_bucket_level_access = true
  force_destroy = true
}
resource "google_storage_batch_operations_job" "job" {
	job_id   = "%s"
	bucket_list {
		buckets  {
			bucket = google_storage_bucket.bucket.name
			prefix_list {
				included_object_prefixes = [
					"objprefix", "prefix2"
				]
			}
		}
	}
	put_object_hold {
		event_based_hold = "SET"
		temporary_hold =  "SET"
	}

	delete_protection = false
}
`, bucketName, jobID)
}

func testAccStorageBatchOperationsJobWithPrefixObjectTemporaryHold(bucketName, jobID string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "%s"
  location = "us-central1"
  uniform_bucket_level_access = true
  force_destroy = true
}
resource "google_storage_batch_operations_job" "job" {
	job_id   = "%s"
	bucket_list {
		buckets  {
			bucket = google_storage_bucket.bucket.name
			prefix_list {
				included_object_prefixes = [
					"objprefix", "prefix2"
				]
			}
		}
	}
	put_object_hold {
		temporary_hold =  "SET"
	}

	delete_protection = false
}
`, bucketName, jobID)
}

func testAccStorageBatchOperationsCreateJobWithManifest(bucketName, jobID string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "%s"
  location = "us-central1"
  uniform_bucket_level_access = true
  force_destroy = true
}
resource "google_storage_batch_operations_job" "job" {
	job_id   = "%s"
	bucket_list {
		buckets  {
			bucket = google_storage_bucket.bucket.name
			manifest {
				manifest_location = "gs://%s/manifest.csv"
			}
		}
	}
	put_metadata {
		custom_metadata = {
			"key"="value"
			"key1"="value1"
		}
		cache_control = "public, max-age=3600"
		content_disposition = "sample"
		content_encoding = "text"
		content_language = "en-us"
		content_type = "application/json"
		custom_time  = "2025-04-30T00:00:00Z"
	}
	delete_protection = false
}
`, bucketName, jobID, bucketName)
}

func testAccStorageBatchOperationsJobs_storageBatchOerationsJobKmsKey(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "%{bucket_name}"
  location = "us-central1"
  uniform_bucket_level_access = true
  force_destroy = true
}

resource "google_storage_bucket_object" "object" {
  name          = "%{object_name}"
  bucket        = google_storage_bucket.bucket.name
  content       = "test-content"
}

data "google_storage_project_service_account" "gcs_account" {
}

resource "google_kms_crypto_key_iam_member" "iam" {
  crypto_key_id = "%{kms_key}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:${data.google_storage_project_service_account.gcs_account.email_address}"
}

resource "google_storage_batch_operations_job" "job" {
	job_id   = "%{job_id}"
	bucket_list {
		buckets  {
			bucket = google_storage_bucket.bucket.name
			prefix_list {
				included_object_prefixes = [
					"objprefix"
				]
			}
		}
	}
	rewrite_object {
		kms_key = "%{kms_key}"
	}

	delete_protection = false
}
`, context)
}
