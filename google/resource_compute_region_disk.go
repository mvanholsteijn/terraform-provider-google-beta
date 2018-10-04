// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/customdiff"
	"github.com/hashicorp/terraform/helper/schema"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
)

func resourceComputeRegionDisk() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionDiskCreate,
		Read:   resourceComputeRegionDiskRead,
		Update: resourceComputeRegionDiskUpdate,
		Delete: resourceComputeRegionDiskDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionDiskImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(300 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},
		CustomizeDiff: customdiff.All(
			customdiff.ForceNewIfChange("size", isDiskShrinkage)),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"replica_zones": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 2,
				MaxItems: 2,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"disk_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"snapshot": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"source_snapshot_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"type": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Default:          "pd-standard",
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_attach_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_detach_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_snapshot_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeRegionDiskCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionDiskDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandComputeRegionDiskLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	nameProp, err := expandComputeRegionDiskName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	sizeGbProp, err := expandComputeRegionDiskSize(d.Get("size"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("size"); !isEmptyValue(reflect.ValueOf(sizeGbProp)) && (ok || !reflect.DeepEqual(v, sizeGbProp)) {
		obj["sizeGb"] = sizeGbProp
	}
	replicaZonesProp, err := expandComputeRegionDiskReplicaZones(d.Get("replica_zones"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("replica_zones"); !isEmptyValue(reflect.ValueOf(replicaZonesProp)) && (ok || !reflect.DeepEqual(v, replicaZonesProp)) {
		obj["replicaZones"] = replicaZonesProp
	}
	typeProp, err := expandComputeRegionDiskType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	regionProp, err := expandComputeRegionDiskRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	diskEncryptionKeyProp, err := expandComputeRegionDiskDiskEncryptionKey(d.Get("disk_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disk_encryption_key"); !isEmptyValue(reflect.ValueOf(diskEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionKeyProp)) {
		obj["diskEncryptionKey"] = diskEncryptionKeyProp
	}
	sourceSnapshotProp, err := expandComputeRegionDiskSnapshot(d.Get("snapshot"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("snapshot"); !isEmptyValue(reflect.ValueOf(sourceSnapshotProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotProp)) {
		obj["sourceSnapshot"] = sourceSnapshotProp
	}
	sourceSnapshotEncryptionKeyProp, err := expandComputeRegionDiskSourceSnapshotEncryptionKey(d.Get("source_snapshot_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_snapshot_encryption_key"); !isEmptyValue(reflect.ValueOf(sourceSnapshotEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotEncryptionKeyProp)) {
		obj["sourceSnapshotEncryptionKey"] = sourceSnapshotEncryptionKeyProp
	}

	obj, err = resourceComputeRegionDiskEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/regions/{{region}}/disks")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionDisk: %#v", obj)
	res, err := sendRequest(config, "POST", url, obj)
	if err != nil {
		return fmt.Errorf("Error creating RegionDisk: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating RegionDisk",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionDisk: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating RegionDisk %q: %#v", d.Id(), res)

	return resourceComputeRegionDiskRead(d, meta)
}

func resourceComputeRegionDiskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/regions/{{region}}/disks/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionDisk %q", d.Id()))
	}

	res, err = resourceComputeRegionDiskDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if err := d.Set("label_fingerprint", flattenComputeRegionDiskLabelFingerprint(res["labelFingerprint"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeRegionDiskCreationTimestamp(res["creationTimestamp"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionDiskDescription(res["description"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("last_attach_timestamp", flattenComputeRegionDiskLastAttachTimestamp(res["lastAttachTimestamp"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("last_detach_timestamp", flattenComputeRegionDiskLastDetachTimestamp(res["lastDetachTimestamp"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("labels", flattenComputeRegionDiskLabels(res["labels"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionDiskName(res["name"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("size", flattenComputeRegionDiskSize(res["sizeGb"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("users", flattenComputeRegionDiskUsers(res["users"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("replica_zones", flattenComputeRegionDiskReplicaZones(res["replicaZones"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("type", flattenComputeRegionDiskType(res["type"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionDiskRegion(res["region"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("disk_encryption_key", flattenComputeRegionDiskDiskEncryptionKey(res["diskEncryptionKey"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("snapshot", flattenComputeRegionDiskSnapshot(res["sourceSnapshot"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("source_snapshot_encryption_key", flattenComputeRegionDiskSourceSnapshotEncryptionKey(res["sourceSnapshotEncryptionKey"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("source_snapshot_id", flattenComputeRegionDiskSourceSnapshotId(res["sourceSnapshotId"])); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionDisk: %s", err)
	}

	return nil
}

func resourceComputeRegionDiskUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	d.Partial(true)

	if d.HasChange("label_fingerprint") || d.HasChange("labels") {
		obj := make(map[string]interface{})
		labelFingerprintProp := d.Get("label_fingerprint")
		obj["labelFingerprint"] = labelFingerprintProp
		labelsProp, err := expandComputeRegionDiskLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}

		url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/regions/{{region}}/disks/{{name}}/setLabels")
		if err != nil {
			return err
		}
		res, err := sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error updating RegionDisk %q: %s", d.Id(), err)
		}

		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating RegionDisk",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("label_fingerprint")
		d.SetPartial("labels")
	}
	if d.HasChange("size") {
		obj := make(map[string]interface{})
		sizeGbProp, err := expandComputeRegionDiskSize(d.Get("size"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("size"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sizeGbProp)) {
			obj["sizeGb"] = sizeGbProp
		}

		url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/regions/{{region}}/disks/{{name}}/resize")
		if err != nil {
			return err
		}
		res, err := sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error updating RegionDisk %q: %s", d.Id(), err)
		}

		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating RegionDisk",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("size")
	}

	d.Partial(false)

	return resourceComputeRegionDiskRead(d, meta)
}

func resourceComputeRegionDiskDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/beta/projects/{{project}}/regions/{{region}}/disks/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	// if disks are attached, they must be detached before the disk can be deleted
	if instances, ok := d.Get("users").([]interface{}); ok {
		type detachArgs struct{ project, zone, instance, deviceName string }
		var detachCalls []detachArgs
		self := d.Get("self_link").(string)
		for _, instance := range instances {
			if !computeDiskUserRegex.MatchString(instance.(string)) {
				return fmt.Errorf("Unknown user %q of disk %q", instance, self)
			}
			matches := computeDiskUserRegex.FindStringSubmatch(instance.(string))
			instanceProject := matches[1]
			instanceZone := matches[2]
			instanceName := matches[3]
			i, err := config.clientCompute.Instances.Get(instanceProject, instanceZone, instanceName).Do()
			if err != nil {
				if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == 404 {
					log.Printf("[WARN] instance %q not found, not bothering to detach disks", instance.(string))
					continue
				}
				return fmt.Errorf("Error retrieving instance %s: %s", instance.(string), err.Error())
			}
			for _, disk := range i.Disks {
				if disk.Source == self {
					detachCalls = append(detachCalls, detachArgs{
						project:    instanceProject,
						zone:       GetResourceNameFromSelfLink(i.Zone),
						instance:   i.Name,
						deviceName: disk.DeviceName,
					})
				}
			}
		}
		for _, call := range detachCalls {
			op, err := config.clientCompute.Instances.DetachDisk(call.project, call.zone, call.instance, call.deviceName).Do()
			if err != nil {
				return fmt.Errorf("Error detaching disk %s from instance %s/%s/%s: %s", call.deviceName, call.project,
					call.zone, call.instance, err.Error())
			}
			err = computeOperationWait(config.clientCompute, op, call.project,
				fmt.Sprintf("Detaching disk from %s/%s/%s", call.project, call.zone, call.instance))
			if err != nil {
				if opErr, ok := err.(ComputeOperationError); ok && len(opErr.Errors) == 1 && opErr.Errors[0].Code == "RESOURCE_NOT_FOUND" {
					log.Printf("[WARN] instance %q was deleted while awaiting detach", call.instance)
					continue
				}
				return err
			}
		}
	}
	log.Printf("[DEBUG] Deleting RegionDisk %q", d.Id())
	res, err := sendRequest(config, "DELETE", url, obj)
	if err != nil {
		return handleNotFoundError(err, d, "RegionDisk")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting RegionDisk",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionDisk %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionDiskImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/disks/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionDiskLabelFingerprint(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskDescription(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskLastAttachTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskLastDetachTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskLabels(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskName(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskSize(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionDiskUsers(v interface{}) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeRegionDiskReplicaZones(v interface{}) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeRegionDiskType(v interface{}) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenComputeRegionDiskRegion(v interface{}) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenComputeRegionDiskDiskEncryptionKey(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeRegionDiskDiskEncryptionKeyRawKey(original["rawKey"])
	transformed["sha256"] =
		flattenComputeRegionDiskDiskEncryptionKeySha256(original["sha256"])
	return []interface{}{transformed}
}
func flattenComputeRegionDiskDiskEncryptionKeyRawKey(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskDiskEncryptionKeySha256(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskSnapshot(v interface{}) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionDiskSourceSnapshotEncryptionKey(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeRegionDiskSourceSnapshotEncryptionKeyRawKey(original["rawKey"])
	transformed["sha256"] =
		flattenComputeRegionDiskSourceSnapshotEncryptionKeySha256(original["sha256"])
	return []interface{}{transformed}
}
func flattenComputeRegionDiskSourceSnapshotEncryptionKeyRawKey(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskSourceSnapshotEncryptionKeySha256(v interface{}) interface{} {
	return v
}

func flattenComputeRegionDiskSourceSnapshotId(v interface{}) interface{} {
	return v
}

func expandComputeRegionDiskDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeRegionDiskName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskSize(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskReplicaZones(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		f, err := parseGlobalFieldValue("zones", raw.(string), "project", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for replica_zones: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeRegionDiskType(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("diskTypes", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for type: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionDiskRegion(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionDiskDiskEncryptionKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeRegionDiskDiskEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeRegionDiskDiskEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	return transformed, nil
}

func expandComputeRegionDiskDiskEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskDiskEncryptionKeySha256(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskSnapshot(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("snapshots", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for snapshot: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionDiskSourceSnapshotEncryptionKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeRegionDiskSourceSnapshotEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeRegionDiskSourceSnapshotEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	return transformed, nil
}

func expandComputeRegionDiskSourceSnapshotEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionDiskSourceSnapshotEncryptionKeySha256(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeRegionDiskEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	// Get the region
	r, err := getRegion(d, config)
	if err != nil {
		return nil, err
	}
	region, err := config.clientCompute.Regions.Get(project, r).Do()
	if err != nil {
		return nil, err
	}

	if v, ok := d.GetOk("type"); ok {
		log.Printf("[DEBUG] Loading disk type: %s", v.(string))
		diskType, err := readRegionDiskType(config, region, project, v.(string))
		if err != nil {
			return nil, fmt.Errorf(
				"Error loading disk type '%s': %s",
				v.(string), err)
		}

		obj["type"] = diskType.SelfLink
	}

	if v, ok := d.GetOk("image"); ok {
		log.Printf("[DEBUG] Resolving image name: %s", v.(string))
		imageUrl, err := resolveImage(config, project, v.(string))
		if err != nil {
			return nil, fmt.Errorf(
				"Error resolving image name '%s': %s",
				v.(string), err)
		}

		obj["sourceImage"] = imageUrl
		log.Printf("[DEBUG] Image name resolved to: %s", imageUrl)
	}

	if v, ok := d.GetOk("snapshot"); ok {
		snapshotName := v.(string)
		match, _ := regexp.MatchString("^https://www.googleapis.com/compute", snapshotName)
		if match {
			obj["sourceSnapshot"] = snapshotName
		} else {
			log.Printf("[DEBUG] Loading snapshot: %s", snapshotName)
			snapshotData, err := config.clientCompute.Snapshots.Get(
				project, snapshotName).Do()

			if err != nil {
				return nil, fmt.Errorf(
					"Error loading snapshot '%s': %s",
					snapshotName, err)
			}
			obj["sourceSnapshot"] = snapshotData.SelfLink
		}
	}

	return obj, nil
}

func resourceComputeRegionDiskDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := res["diskEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("disk_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]
		res["diskEncryptionKey"] = transformed
	}

	if v, ok := res["sourceImageEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("source_image_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]
		res["sourceImageEncryptionKey"] = transformed
	}

	if v, ok := res["sourceSnapshotEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("source_snapshot_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]
		res["sourceSnapshotEncryptionKey"] = transformed
	}

	return res, nil
}
