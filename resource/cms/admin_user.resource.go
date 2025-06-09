package cms

// CreateAdminUserRequest is a request for creating admin user
type CreateAdminUserRequest struct {
	Email  string `json:"email" binding:"required,email"`
	Name   string `json:"name" binding:"required,min=4,max=100"`
	PIN    string `json:"pin" form:"pin" binding:"required,len=6"`
	RoleID string `json:"role_id" binding:"required,uuid4"`
}

// UpdateAdminUserRequest is a request for update admin user
type UpdateAdminUserRequest struct {
	Email  string `json:"email" binding:"required,email"`
	Name   string `json:"name" binding:"required,min=4,max=100"`
	RoleID string `json:"role_id" binding:"required,uuid4"`
	PIN    string `json:"pin" form:"pin" binding:"omitempty,len=6,numeric"`
}
