package entity

import (
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
	Email    string    `json:"email"`
	Password string    `json:"password"`
	UserRole *UserRole `foreignKey:"ID" associationForeignKey:"UserID"`
	Gender   string    `json:"gender"`
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
	email string,
	password string,
	gender string,
	createdBy string,
) *User {
	if password != "" {
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		password = string(passwordHash)
	}

	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		Gender:    gender,
		Auditable: NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *User) MapUpdateFrom(from *User) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"username": model.Username,
			"email":    model.Email,
			"gender":   model.Gender,
		}
	}

	mapped := make(map[string]interface{})

	if model.Username != from.Username {
		mapped["username"] = from.Username
	}

	if model.Email != from.Email {
		mapped["email"] = from.Email
	}

	if model.Gender != from.Gender {
		mapped["Gender"] = from.Gender
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
