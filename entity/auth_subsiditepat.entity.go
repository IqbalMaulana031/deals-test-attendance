package entity

// TokenSubsidiTepat represents an access token
type TokenSubsidiTepat struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	CreatedAt    int64  `json:"created_at"`
}

// TokenInfoSubsidiTepat represents token information
type TokenInfoSubsidiTepat struct {
	ExpiryIn  int    `json:"expiryIn"`
	CreatedAt int64  `json:"createdAt"`
	Name      string `json:"name"`
}
