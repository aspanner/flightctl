package model

import (
	"encoding/json"
	"time"

	api "github.com/flightctl/flightctl/api/v1alpha1"
	"github.com/flightctl/flightctl/internal/util"
)

var (
	DeviceAPI      = "v1alpha1"
	DeviceKind     = "Device"
	DeviceListKind = "DeviceList"
)

type Device struct {
	Resource

	// The desired state, stored as opaque JSON object.
	Spec *JSONField[api.DeviceSpec]

	// The last reported state, stored as opaque JSON object.
	Status *JSONField[api.DeviceStatus]
}

type DeviceList []Device

func (d Device) String() string {
	val, _ := json.Marshal(d)
	return string(val)
}

func NewDeviceFromApiResource(resource *api.Device) *Device {
	if resource == nil || resource.Metadata.Name == nil {
		return &Device{}
	}

	var status api.DeviceStatus
	if resource.Status != nil {
		status = *resource.Status
	}
	return &Device{
		Resource: Resource{
			Name: *resource.Metadata.Name,
		},
		Spec:   MakeJSONField(resource.Spec),
		Status: MakeJSONField(status),
	}
}

func (d *Device) ToApiResource() api.Device {
	if d == nil {
		return api.Device{}
	}

	var status api.DeviceStatus
	if d.Status != nil {
		status = d.Status.Data
	}
	return api.Device{
		ApiVersion: DeviceAPI,
		Kind:       DeviceKind,
		Metadata: api.ObjectMeta{
			Name:              &d.Name,
			CreationTimestamp: util.StrToPtr(d.CreatedAt.UTC().Format(time.RFC3339)),
		},
		Spec:   d.Spec.Data,
		Status: &status,
	}
}

func (dl DeviceList) ToApiResource() api.DeviceList {
	if dl == nil {
		return api.DeviceList{
			ApiVersion: DeviceAPI,
			Kind:       DeviceListKind,
			Items:      []api.Device{},
		}
	}

	deviceList := make([]api.Device, len(dl))
	for i, device := range dl {
		deviceList[i] = device.ToApiResource()
	}
	return api.DeviceList{
		ApiVersion: DeviceAPI,
		Kind:       DeviceListKind,
		Items:      deviceList,
	}
}
