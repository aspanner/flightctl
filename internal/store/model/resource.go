package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Resource struct {
	// Uniquely identifies the tenant the resource belongs to.
	// Assigned by IAM. Immutable.
	OrgID uuid.UUID `gorm:"type:uuid;primary_key;"`

	// Uniquely identifies the resource within a tenant and schema.
	// Depending on the schema (kind), assigned by the device management system or the crypto identity of the device (public key). Immutable.
	// This may become a URN later, so it's important API users treat this as an opaque handle.
	Name string `gorm:"primary_key;"`

	// User-defined name, if non-null used in the UI as a more human-friendly alias to the resource ID.
	// DisplayName string

	// The "kind/name" of the resource owner of this resource.
	// Owner string

	// User-defined labels, used to select resources in queries.
	// Labels StringArray

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (r *Resource) BeforeCreate(tx *gorm.DB) error {
	if len(r.Name) == 0 {
		r.Name = uuid.New().String()
	}
	return nil
}
