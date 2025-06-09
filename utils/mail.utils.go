package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"starter-go-gin/common/errors"
)

const (
	two = 2
)

// EmailPayload is the payload for sending email
type EmailPayload struct {
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

// ConstructEmailPayload is a function to construct an email payload
func ConstructEmailPayload(templatePath, receiver, subject, category string, data map[string]interface{}) (*EmailPayload, error) {
	var body bytes.Buffer

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println(fmt.Errorf("failed to load email template: %w", err))
		return nil, errors.ErrInternalServerError.Error()
	}

	err = t.Execute(&body, data)
	if err != nil {
		log.Println(fmt.Errorf("failed to execute email data: %w", err))
		return nil, errors.ErrInternalServerError.Error()
	}

	emailPayload := &EmailPayload{
		To:       receiver,
		Subject:  subject,
		Content:  body.String(),
		Category: category,
	}

	return emailPayload, nil
}

func GetDomainSubstring(email string) (string, error) {
	// Split the email address by "@"
	parts := strings.Split(email, "@")
	if len(parts) != two {
		return "", fmt.Errorf("invalid email address")
	}

	// Get the domain part (after "@")
	domainPart := parts[1]

	// Split the domain part by "."
	domainParts := strings.Split(domainPart, ".")
	if len(domainParts) < two {
		return "", fmt.Errorf("invalid domain in email address")
	}

	// Get the main domain name (usually the second last part)
	mainDomain := domainParts[0]

	return mainDomain, nil
}
