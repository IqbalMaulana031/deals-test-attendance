package fcm

import (
	"context"
	"log"

	"google.golang.org/api/fcm/v1"
	"google.golang.org/api/option"

	"starter-go-gin/config"
)

// FCM is an struct for firebase cloud messaging SDK
type FCM struct {
	cfg     config.Config
	service *fcm.Service
}

// NewFCM initiate FCM SDK
func NewFCM(cfg config.Config) *FCM {
	ctx := context.Background()
	service, err := fcm.NewService(ctx, option.WithScopes(fcm.FirebaseMessagingScope))
	if err != nil {
		log.Fatalf("Failed to create fcm service: %v", err)
	}

	return &FCM{
		cfg:     cfg,
		service: service,
	}
}

// SendNotification send notification to device
func (f *FCM) SendNotification(token string, title string, body string, data map[string]string) error {
	msg := &fcm.Message{
		Notification: &fcm.Notification{
			Title: title,
			Body:  body,
		},
		Data:  data,
		Token: token,
		Android: &fcm.AndroidConfig{
			Priority: "HIGH",
		},
	}

	sendMessageRequest := &fcm.SendMessageRequest{
		Message: msg,
	}

	res, err := f.service.Projects.Messages.Send("projects/"+f.cfg.Google.ProjectID, sendMessageRequest).Do()
	print(res)
	return err
}
