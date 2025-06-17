package entity

import (
	"starter-go-gin/utils"
	"time"

	"github.com/google/uuid"
)

const (
	roleTableName = "auth.roles"
)

// Role defines table role
type Role struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
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
	createdBy string,
) *Role {
	return &Role{
		ID:        id,
		Name:      name,
		Auditable: NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *Role) MapUpdateFrom(from *Role) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":       model.Name,
			"updated_by": model.CreatedBy,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	mapped["updated_at"] = utils.AddSevenHours(time.Now())
	return &mapped
}
