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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/documentai/ProcessorDefaultVersion.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package documentai

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

func ResourceDocumentAIProcessorDefaultVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceDocumentAIProcessorDefaultVersionCreate,
		Read:   resourceDocumentAIProcessorDefaultVersionRead,
		Delete: resourceDocumentAIProcessorDefaultVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDocumentAIProcessorDefaultVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"processor": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The processor to set the version on.`,
			},
			"version": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.ProjectNumberDiffSuppress,
				Description: `The version to set. Using 'stable' or 'rc' will cause the API to return the latest version in that release channel.
Apply 'lifecycle.ignore_changes' to the 'version' field to suppress this diff.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDocumentAIProcessorDefaultVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	defaultProcessorVersionProp, err := expandDocumentAIProcessorDefaultVersionVersion(d.Get("version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultProcessorVersionProp)) && (ok || !reflect.DeepEqual(v, defaultProcessorVersionProp)) {
		obj["defaultProcessorVersion"] = defaultProcessorVersionProp
	}
	processorProp, err := expandDocumentAIProcessorDefaultVersionProcessor(d.Get("processor"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("processor"); !tpgresource.IsEmptyValue(reflect.ValueOf(processorProp)) && (ok || !reflect.DeepEqual(v, processorProp)) {
		obj["processor"] = processorProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DocumentAIBasePath}}{{processor}}:setDefaultProcessorVersion")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ProcessorDefaultVersion: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	if strings.Contains(url, "https://-") {
		location := tpgresource.GetRegionFromRegionalSelfLink(url)
		url = strings.TrimPrefix(url, "https://")
		url = "https://" + location + url
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
		return fmt.Errorf("Error creating ProcessorDefaultVersion: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{processor}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ProcessorDefaultVersion %q: %#v", d.Id(), res)

	return resourceDocumentAIProcessorDefaultVersionRead(d, meta)
}

func resourceDocumentAIProcessorDefaultVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DocumentAIBasePath}}{{processor}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	if strings.Contains(url, "https://-") {
		location := tpgresource.GetRegionFromRegionalSelfLink(url)
		url = strings.TrimPrefix(url, "https://")
		url = "https://" + location + url
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DocumentAIProcessorDefaultVersion %q", d.Id()))
	}

	if err := d.Set("version", flattenDocumentAIProcessorDefaultVersionVersion(res["defaultProcessorVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProcessorDefaultVersion: %s", err)
	}

	return nil
}

func resourceDocumentAIProcessorDefaultVersionDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] DocumentAI ProcessorDefaultVersion resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceDocumentAIProcessorDefaultVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<processor>.+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{processor}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDocumentAIProcessorDefaultVersionVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDocumentAIProcessorDefaultVersionVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDocumentAIProcessorDefaultVersionProcessor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
