package utils

import (
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// GetClientIPHelper gets the client IP using a mixture of techniques.
// This is how it is with golang at the moment.
func GetClientIPHelper(req *http.Request) (string, error) {
	// Try Request Headers (X-Forwarder). Client could be behind a Proxy
	ip, err := getClientIPByHeaders(req)
	if err == nil {
		log.Printf("debug: Found IP using Request Headers sniffing. ip: %v", ip)
		if ip != "::1" {
			return ip, nil
		}
	}

	err = errors.New("error: Could not find clients IP address")
	return "", err
}

// getClientIPByHeaders tries to get directly from the Request Headers.
// This is only way when the client is behind a Proxy.
func getClientIPByHeaders(req *http.Request) (string, error) {
	// Client could be behid a Proxy, so Try Request Headers (X-Forwarder)
	ipSlice := []string{}

	ipSlice = append(ipSlice, req.Header.Get("X-Forwarded-For"))
	ipSlice = append(ipSlice, req.Header.Get("x-forwarded-for"))
	ipSlice = append(ipSlice, req.Header.Get("X-FORWARDED-FOR"))

	for _, v := range ipSlice {
		log.Printf("debug: client request header check gives ip: %v", v)
		if v != "" {
			return v, nil
		}
	}
	err := errors.New("error: Could not find clients IP address from the Request Headers")
	return "", err
}

// FindString is a helper function to find a string in a slice of strings
func FindString(collectString []string, findString string) bool {
	for i := range collectString {
		if collectString[i] == findString {
			return true
		}
	}
	return false
}

// SetZeroPhoneNumber sets phone number to 0 if it starts with 62
func SetZeroPhoneNumber(phoneNumber string) (string, error) {
	// check if phone number starts with 0 or 62
	if phoneNumber[0] != '0' && phoneNumber[0:2] != "62" {
		return "", errors.New("error: Phone number is not valid")
	}

	if phoneNumber[0:2] == "62" {
		phoneNumber = "0" + phoneNumber[2:]
	}

	return phoneNumber, nil
}
