package entity

import (
	"starter-go-gin/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	userTableName = "auth.users"
)

// User is a model for user
type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Sallary  int       `json:"salary"`
	RoleID   uuid.UUID `json:"role_id"`
	Role     Role      `foreignKey:"ID" associationForeignKey:"RoleID" json:"role"`
	Auditable
}

// TableName specifies table name
func (model *User) TableName() string {
	return userTableName
}

// NewUser is a constructor for user
func NewUser(
	id uuid.UUID,
	username string,
	password string,
	Sallary int,
	RoleID uuid.UUID,
	createdBy string,
) *User {
	if password != "" {
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		password = string(passwordHash)
	}

	return &User{
		ID:        id,
		Username:  username,
		Password:  password,
		Sallary:   Sallary,
		RoleID:    RoleID,
		Auditable: NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *User) MapUpdateFrom(from *User) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"username":   model.Username,
			"salary":     model.Sallary,
			"role_id":    model.RoleID,
			"updated_by": model.CreatedBy,
			"updated_at": utils.AddSevenHours(time.Now()),
		}
	}

	mapped := make(map[string]interface{})

	if model.Username != from.Username {
		mapped["username"] = from.Username
	}

	if model.Sallary != from.Sallary {
		mapped["salary"] = from.Sallary
	}

	if model.RoleID != from.RoleID {
		mapped["role_id"] = from.RoleID
	}

	mapped["updated_at"] = utils.AddSevenHours(time.Now())
	return &mapped
}
