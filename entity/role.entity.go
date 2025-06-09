package entity

import (
	"github.com/google/uuid"
)

const (
	roleTableName = "auth.roles"
)

// Role defines table role
type Role struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Label string    `json:"label"`
	Type  string    `json:"type"`
	Auditable
}

// TableName specifies table name
func (model *Role) TableName() string {
	return roleTableName
}

// NewRole creates new role entity
func NewRole(
	id uuid.UUID,
	name string,
	label string,
	roleType string,
	createdBy string,
) *Role {
	return &Role{
		ID:        id,
		Name:      name,
		Label:     label,
		Type:      roleType,
		Auditable: NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *Role) MapUpdateFrom(from *Role) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":  model.Name,
			"label": model.Label,
			"type":  model.Type,
		}
	}

	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	if model.Label != from.Label {
		mapped["label"] = from.Label
	}

	if model.Type != from.Type {
		mapped["type"] = from.Type
	}

	mapped["updated_at"] = from.UpdatedAt

	return &mapped
}
