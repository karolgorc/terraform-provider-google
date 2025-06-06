// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/oslogin/SSHPublicKey.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package oslogin

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceOSLoginSSHPublicKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceOSLoginSSHPublicKeyCreate,
		Read:   resourceOSLoginSSHPublicKeyRead,
		Update: resourceOSLoginSSHPublicKeyUpdate,
		Delete: resourceOSLoginSSHPublicKeyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceOSLoginSSHPublicKeyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Public key text in SSH format, defined by RFC4253 section 6.6.`,
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The user email.`,
			},
			"expiration_time_usec": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An expiration time in microseconds since epoch.`,
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The project ID of the Google Cloud Platform project.`,
			},
			"fingerprint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The SHA-256 fingerprint of the SSH public key.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceOSLoginSSHPublicKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	keyProp, err := expandOSLoginSSHPublicKeyKey(d.Get("key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("key"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyProp)) && (ok || !reflect.DeepEqual(v, keyProp)) {
		obj["key"] = keyProp
	}
	expirationTimeUsecProp, err := expandOSLoginSSHPublicKeyExpirationTimeUsec(d.Get("expiration_time_usec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_time_usec"); !tpgresource.IsEmptyValue(reflect.ValueOf(expirationTimeUsecProp)) && (ok || !reflect.DeepEqual(v, expirationTimeUsecProp)) {
		obj["expirationTimeUsec"] = expirationTimeUsecProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{OSLoginBasePath}}users/{{user}}:importSshPublicKey")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SSHPublicKey: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	// Don't use `GetProject()` because we only want to set the project in the URL
	// if the user set it explicitly on the resource.
	if p, ok := d.GetOk("project"); ok {
		url, err = transport_tpg.AddQueryParams(url, map[string]string{"projectId": p.(string)})
		if err != nil {
			return err
		}
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating SSHPublicKey: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	err = resourceOSLoginSSHPublicKeyPostCreateSetComputedFields(d, meta, res)
	if err != nil {
		return fmt.Errorf("setting computed ID format fields: %w", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "users/{{user}}/sshPublicKeys/{{fingerprint}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	loginProfile, ok := res["loginProfile"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}

	// `fingerprint` is autogenerated from the api so needs to be set post-create
	sshPublicKeys := loginProfile.(map[string]interface{})["sshPublicKeys"]
	for _, sshPublicKey := range sshPublicKeys.(map[string]interface{}) {
		if sshPublicKey.(map[string]interface{})["key"].(string) == d.Get("key") {
			if err := d.Set("fingerprint", sshPublicKey.(map[string]interface{})["fingerprint"].(string)); err != nil {
				return fmt.Errorf("Error setting fingerprint: %s", err)
			}
			break
		}
	}

	// Store the ID now
	id, err = tpgresource.ReplaceVars(d, config, "users/{{user}}/sshPublicKeys/{{fingerprint}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Wait 10s for write propagation; should be usually done in ~7s max.
	time.Sleep(10 * time.Second)

	log.Printf("[DEBUG] Finished creating SSHPublicKey %q: %#v", d.Id(), res)

	return resourceOSLoginSSHPublicKeyRead(d, meta)
}

func resourceOSLoginSSHPublicKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{OSLoginBasePath}}users/{{user}}/sshPublicKeys/{{fingerprint}}/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("OSLoginSSHPublicKey %q", d.Id()))
	}

	if err := d.Set("key", flattenOSLoginSSHPublicKeyKey(res["key"], d, config)); err != nil {
		return fmt.Errorf("Error reading SSHPublicKey: %s", err)
	}
	if err := d.Set("expiration_time_usec", flattenOSLoginSSHPublicKeyExpirationTimeUsec(res["expirationTimeUsec"], d, config)); err != nil {
		return fmt.Errorf("Error reading SSHPublicKey: %s", err)
	}
	if err := d.Set("fingerprint", flattenOSLoginSSHPublicKeyFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading SSHPublicKey: %s", err)
	}

	return nil
}

func resourceOSLoginSSHPublicKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	expirationTimeUsecProp, err := expandOSLoginSSHPublicKeyExpirationTimeUsec(d.Get("expiration_time_usec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_time_usec"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, expirationTimeUsecProp)) {
		obj["expirationTimeUsec"] = expirationTimeUsecProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{OSLoginBasePath}}users/{{user}}/sshPublicKeys/{{fingerprint}}/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SSHPublicKey %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("expiration_time_usec") {
		updateMask = append(updateMask, "expirationTimeUsec")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating SSHPublicKey %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating SSHPublicKey %q: %#v", d.Id(), res)
		}

	}

	return resourceOSLoginSSHPublicKeyRead(d, meta)
}

func resourceOSLoginSSHPublicKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{OSLoginBasePath}}users/{{user}}/sshPublicKeys/{{fingerprint}}/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting SSHPublicKey %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "SSHPublicKey")
	}

	log.Printf("[DEBUG] Finished deleting SSHPublicKey %q: %#v", d.Id(), res)
	return nil
}

func resourceOSLoginSSHPublicKeyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^users/(?P<user>[^/]+)/sshPublicKeys/(?P<fingerprint>[^/]+)$",
		"^(?P<user>[^/]+)/(?P<fingerprint>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "users/{{user}}/sshPublicKeys/{{fingerprint}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenOSLoginSSHPublicKeyKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenOSLoginSSHPublicKeyExpirationTimeUsec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenOSLoginSSHPublicKeyFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandOSLoginSSHPublicKeyKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOSLoginSSHPublicKeyExpirationTimeUsec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceOSLoginSSHPublicKeyPostCreateSetComputedFields(d *schema.ResourceData, meta interface{}, res map[string]interface{}) error {
	config := meta.(*transport_tpg.Config)
	if err := d.Set("fingerprint", flattenOSLoginSSHPublicKeyFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "fingerprint": %s`, err)
	}
	return nil
}
