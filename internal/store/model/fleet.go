package model

import (
	"encoding/json"
	"time"

	api "github.com/flightctl/flightctl/api/v1alpha1"
	"github.com/flightctl/flightctl/internal/util"
)

var (
	FleetAPI      = "v1alpha1"
	FleetKind     = "Fleet"
	FleetListKind = "FleetList"
)

type Fleet struct {
	Resource

	// The desired state, stored as opaque JSON object.
	Spec *JSONField[api.FleetSpec]

	// The last reported state, stored as opaque JSON object.
	Status *JSONField[api.FleetStatus]
}

type FleetList []Fleet

func (d Fleet) String() string {
	val, _ := json.Marshal(d)
	return string(val)
}

func NewFleetFromApiResource(resource *api.Fleet) *Fleet {
	if resource == nil || resource.Metadata.Name == nil {
		return &Fleet{}
	}

	var status api.FleetStatus
	if resource.Status != nil {
		status = *resource.Status
	}
	return &Fleet{
		Resource: Resource{
			Name: *resource.Metadata.Name,
		},
		Spec:   MakeJSONField(resource.Spec),
		Status: MakeJSONField(status),
	}
}

func (f *Fleet) ToApiResource() api.Fleet {
	if f == nil {
		return api.Fleet{}
	}

	var status api.FleetStatus
	if f.Status != nil {
		status = f.Status.Data
	}
	return api.Fleet{
		ApiVersion: FleetAPI,
		Kind:       FleetKind,
		Metadata: api.ObjectMeta{
			Name:              &f.Name,
			CreationTimestamp: util.StrToPtr(f.CreatedAt.UTC().Format(time.RFC3339)),
		},
		Spec:   f.Spec.Data,
		Status: &status,
	}
}

func (dl FleetList) ToApiResource() api.FleetList {
	if dl == nil {
		return api.FleetList{
			ApiVersion: FleetAPI,
			Kind:       FleetListKind,
			Items:      []api.Fleet{},
		}
	}

	fleetList := make([]api.Fleet, len(dl))
	for i, fleet := range dl {
		fleetList[i] = fleet.ToApiResource()
	}
	return api.FleetList{
		ApiVersion: FleetAPI,
		Kind:       FleetListKind,
		Items:      fleetList,
	}
}
