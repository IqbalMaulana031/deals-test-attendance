package interfaces

// FCMUseCase is an interface for firebase cloud messaging SDK
type FCMUseCase interface {
	SendNotification(token string, title string, body string, data map[string]string) error
}
