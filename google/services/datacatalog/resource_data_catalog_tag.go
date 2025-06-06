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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/datacatalog/Tag.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package datacatalog

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceDataCatalogTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataCatalogTagCreate,
		Read:   resourceDataCatalogTagRead,
		Update: resourceDataCatalogTagUpdate,
		Delete: resourceDataCatalogTagDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataCatalogTagImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		DeprecationMessage: "`google_data_catalog_tag` is deprecated and will be removed in a future major release. For steps to transition your Data Catalog users, workloads, and content to Dataplex Catalog, see https://cloud.google.com/dataplex/docs/transition-to-dataplex-catalog.",

		Schema: map[string]*schema.Schema{
			"fields": {
				Type:     schema.TypeSet,
				Required: true,
				Description: `This maps the ID of a tag field to the value of and additional information about that field.
Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"bool_value": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Holds the value for a tag field with boolean type.`,
						},
						"double_value": {
							Type:        schema.TypeFloat,
							Optional:    true,
							Description: `Holds the value for a tag field with double type.`,
						},
						"enum_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The display name of the enum value.`,
						},

						"string_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Holds the value for a tag field with string type.`,
						},
						"timestamp_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Holds the value for a tag field with timestamp type.`,
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The display name of this field`,
						},
						"order": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: `The order of this field with respect to other fields in this tag. For example, a higher value can indicate
a more important field. The value can be negative. Multiple fields can have the same order, and field orders
within a tag do not have to be sequential.`,
						},
					},
				},
			},
			"template": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource name of the tag template that this tag uses. Example:
projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
This field cannot be modified after creation.`,
			},
			"column": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
individual column based on that schema.

For attaching a tag to a nested column, use '.' to separate the column names. Example:
'outer_column.inner_column'`,
			},
			"parent": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
all entries in that group.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the tag in URL format. Example:
projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id}
where tag_id is a system-generated identifier. Note that this Tag may not actually be stored in the location in this name.`,
			},
			"template_displayname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The display name of the tag template.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDataCatalogTagCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	templateProp, err := expandNestedDataCatalogTagTemplate(d.Get("template"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("template"); !tpgresource.IsEmptyValue(reflect.ValueOf(templateProp)) && (ok || !reflect.DeepEqual(v, templateProp)) {
		obj["template"] = templateProp
	}
	fieldsProp, err := expandNestedDataCatalogTagFields(d.Get("fields"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fields"); !tpgresource.IsEmptyValue(reflect.ValueOf(fieldsProp)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}
	columnProp, err := expandNestedDataCatalogTagColumn(d.Get("column"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("column"); !tpgresource.IsEmptyValue(reflect.ValueOf(columnProp)) && (ok || !reflect.DeepEqual(v, columnProp)) {
		obj["column"] = columnProp
	}

	obj, err = resourceDataCatalogTagEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{parent}}/tags")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Tag: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
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
		return fmt.Errorf("Error creating Tag: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	err = resourceDataCatalogTagPostCreateSetComputedFields(d, meta, res)
	if err != nil {
		return fmt.Errorf("setting computed ID format fields: %w", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Tag %q: %#v", d.Id(), res)

	return resourceDataCatalogTagRead(d, meta)
}

func resourceDataCatalogTagRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{parent}}/tags?pageSize=1000")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataCatalogTag %q", d.Id()))
	}

	res, err = flattenNestedDataCatalogTag(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing DataCatalogTag because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenNestedDataCatalogTagName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tag: %s", err)
	}
	if err := d.Set("template", flattenNestedDataCatalogTagTemplate(res["template"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tag: %s", err)
	}
	if err := d.Set("template_displayname", flattenNestedDataCatalogTagTemplateDisplayname(res["templateDisplayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tag: %s", err)
	}
	if err := d.Set("fields", flattenNestedDataCatalogTagFields(res["fields"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tag: %s", err)
	}
	if err := d.Set("column", flattenNestedDataCatalogTagColumn(res["column"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tag: %s", err)
	}

	return nil
}

func resourceDataCatalogTagUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	fieldsProp, err := expandNestedDataCatalogTagFields(d.Get("fields"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fields"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}
	columnProp, err := expandNestedDataCatalogTagColumn(d.Get("column"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("column"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, columnProp)) {
		obj["column"] = columnProp
	}

	obj, err = resourceDataCatalogTagEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Tag %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("fields") {
		updateMask = append(updateMask, "fields")
	}

	if d.HasChange("column") {
		updateMask = append(updateMask, "column")
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
			return fmt.Errorf("Error updating Tag %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Tag %q: %#v", d.Id(), res)
		}

	}

	return resourceDataCatalogTagRead(d, meta)
}

func resourceDataCatalogTagDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Tag %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Tag")
	}

	log.Printf("[DEBUG] Finished deleting Tag %q: %#v", d.Id(), res)
	return nil
}

func resourceDataCatalogTagImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	egRegex := regexp.MustCompile("(.+)/tags")

	parts := egRegex.FindStringSubmatch(name)
	if len(parts) != 2 {
		return nil, fmt.Errorf("entry name does not fit the format %s", egRegex)
	}

	if err := d.Set("parent", parts[1]); err != nil {
		return nil, fmt.Errorf("Error setting parent: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenNestedDataCatalogTagName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedDataCatalogTagTemplate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedDataCatalogTagTemplateDisplayname(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedDataCatalogTagFields(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"field_name":      k,
			"display_name":    flattenNestedDataCatalogTagFieldsDisplayName(original["display_name"], d, config),
			"order":           flattenNestedDataCatalogTagFieldsOrder(original["order"], d, config),
			"double_value":    flattenNestedDataCatalogTagFieldsDoubleValue(original["doubleValue"], d, config),
			"string_value":    flattenNestedDataCatalogTagFieldsStringValue(original["stringValue"], d, config),
			"bool_value":      flattenNestedDataCatalogTagFieldsBoolValue(original["boolValue"], d, config),
			"timestamp_value": flattenNestedDataCatalogTagFieldsTimestampValue(original["timestampValue"], d, config),
			"enum_value":      flattenNestedDataCatalogTagFieldsEnumValue(original["enumValue"], d, config),
		})
	}
	return transformed
}
func flattenNestedDataCatalogTagFieldsDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedDataCatalogTagFieldsOrder(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenNestedDataCatalogTagFieldsDoubleValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedDataCatalogTagFieldsStringValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedDataCatalogTagFieldsBoolValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedDataCatalogTagFieldsTimestampValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedDataCatalogTagFieldsEnumValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}

	return v.(map[string]interface{})["displayName"]
}

func flattenNestedDataCatalogTagColumn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedDataCatalogTagTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedDataCatalogTagFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandNestedDataCatalogTagFieldsDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["display_name"] = transformedDisplayName
		}

		transformedOrder, err := expandNestedDataCatalogTagFieldsOrder(original["order"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOrder); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["order"] = transformedOrder
		}

		transformedDoubleValue, err := expandNestedDataCatalogTagFieldsDoubleValue(original["double_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDoubleValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["doubleValue"] = transformedDoubleValue
		}

		transformedStringValue, err := expandNestedDataCatalogTagFieldsStringValue(original["string_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStringValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["stringValue"] = transformedStringValue
		}

		transformedBoolValue, err := expandNestedDataCatalogTagFieldsBoolValue(original["bool_value"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["boolValue"] = transformedBoolValue
		}

		transformedTimestampValue, err := expandNestedDataCatalogTagFieldsTimestampValue(original["timestamp_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTimestampValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["timestampValue"] = transformedTimestampValue
		}

		transformedEnumValue, err := expandNestedDataCatalogTagFieldsEnumValue(original["enum_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnumValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["enumValue"] = transformedEnumValue
		}

		transformedFieldName, err := tpgresource.ExpandString(original["field_name"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedFieldName] = transformed
	}
	return m, nil
}

func expandNestedDataCatalogTagFieldsDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedDataCatalogTagFieldsOrder(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedDataCatalogTagFieldsDoubleValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedDataCatalogTagFieldsStringValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedDataCatalogTagFieldsBoolValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedDataCatalogTagFieldsTimestampValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedDataCatalogTagFieldsEnumValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	// we flattened the original["enum_value"]["display_name"] object to be just original["enum_value"] so here,
	// v is the value we want from the config
	transformed := make(map[string]interface{})
	if val := reflect.ValueOf(v); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["displayName"] = v
	}

	return transformed, nil
}

func expandNestedDataCatalogTagColumn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceDataCatalogTagEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	if obj["fields"] != nil {
		// IsEmptyValue() does not work for a boolean as it shows
		// false when it is 'empty'. Filter boolValue here based on
		// the rule api does not take more than 1 'value'
		fields := obj["fields"].(map[string]interface{})
		for _, elements := range fields {
			values := elements.(map[string]interface{})
			if len(values) > 1 {
				for val := range values {
					if val == "boolValue" {
						delete(values, "boolValue")
					}
				}
			}
		}
	}
	return obj, nil
}

func flattenNestedDataCatalogTag(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["tags"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value tags. Actual value: %v", v)
	}

	_, item, err := resourceDataCatalogTagFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceDataCatalogTagFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName := d.Get("name")
	expectedFlattenedName := flattenNestedDataCatalogTagName(expectedName, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemName := flattenNestedDataCatalogTagName(item["name"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemName)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedName))) && !reflect.DeepEqual(itemName, expectedFlattenedName) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemName, expectedFlattenedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceDataCatalogTagPostCreateSetComputedFields(d *schema.ResourceData, meta interface{}, res map[string]interface{}) error {
	config := meta.(*transport_tpg.Config)
	if _, ok := res["tags"]; ok {
		res, err := flattenNestedDataCatalogTag(d, meta, res)
		if err != nil {
			return fmt.Errorf("Error getting nested object from operation response: %s", err)
		}
		if res == nil {
			// Object isn't there any more - remove it from the state.
			return fmt.Errorf("Error decoding response from operation, could not find nested object")
		}
	}
	if err := d.Set("name", flattenNestedDataCatalogTagName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}
	return nil
}
