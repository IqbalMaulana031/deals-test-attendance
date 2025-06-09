package resource

// ReminderChangePINResponse is a struct for reminder change pin response
type ReminderChangePINResponse struct {
	Redirect       string `json:"redirect"`
	RefreshmentPIN bool   `json:"refreshment_pin"`
}

// ReminderChangePIN is a constructor for ReminderChangePIN
type ReminderChangePIN struct {
	UserID  string `json:"user_id"`
	Expired string `json:"expired"`
}

// ValidateOTPEmailRequest is a request for ValidateOTPEmail
type ValidateOTPEmailRequest struct {
	Type string `json:"type" binding:"required"`
	OTP  string `json:"otp" validate:"required,numeric,len=4"`
}

// RequestOTPNewEmailRequest is a constructor for RequestOTPNewEmailRequest
type RequestOTPNewEmailRequest struct {
	Email string `json:"email" binding:"required,email,max=50"`
}

// GetProfileByTokenRequest is a request for get profile by token
type GetProfileByTokenRequest struct {
	Token string `uri:"token" binding:"required"`
}

// GetProfileByTokenResponse is a response for get profile by token
type GetProfileByTokenResponse struct {
	Email string `json:"email"`
}

// LoginPDSRequest is a request for login pds
type LoginPDSRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	PIN      string `form:"pin" json:"pin" binding:"required,numeric,len=6"`
}

// SSOPDSRequest is a request for sso pds
type SSOPDSRequest struct {
	SessionKey string `json:"session_key" form:"session_key" binding:"required"`
}

// LoginRequest is a struct for login request
type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" form:"username_or_email" binding:"required,max=50"`
	Password        string `json:"password" form:"password" binding:"required"`
}

// LoginRequestAdmin is a struct for login request
type LoginRequestAdmin struct {
	Email    string `json:"email" form:"email" binding:"required,max=50"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LogoutRequest is a struct for logout request
type LogoutRequest struct {
	DeviceID string `json:"device_id" form:"device_id" binding:"omitempty,max=200"`
}

// LoginCMSRequest is a struct for login request
type LoginCMSRequest struct {
	Email   string `json:"email" form:"email" binding:"required,email,max=50"`
	PIN     string `form:"pin" json:"pin" binding:"required"`
	Captcha string `json:"captcha" form:"captcha" binding:"required"`
}

// LoginResponse is a struct for login response
type LoginResponse struct {
	Token string `json:"token"`
}

// NewLoginResponse is a constructor for LoginResponse
func NewLoginResponse(token string) *LoginResponse {
	return &LoginResponse{Token: token}
}

// DeviceRequest is a struct for device request
type DeviceRequest struct {
	DeviceToken string `json:"device_token" binding:"max=255"`
	DeviceType  string `json:"device_type" binding:"max=10"`
	DeviceID    string `json:"device_id,omitempty" binding:"max=50"`
}

// RegisterRequest is a struct for register request
type RegisterRequest struct {
	Username          string `json:"username" binding:"required" form:"username"`
	Email             string `json:"email" binding:"required,email" form:"email"`
	Password          string `json:"password"`
	DuplicatePassword string `json:"duplicate_password" binding:"required" form:"duplicate_password"`
	Gender            string `json:"gender" binding:"omitempty"`
}
