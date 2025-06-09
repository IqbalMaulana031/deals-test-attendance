package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	five = 5
)

// CallerHeader define struct fir api call header
type CallerHeader struct {
	Key   string
	Value string
}

func CallAPI(httpMethod, fullURL string, headers []CallerHeader, payload interface{}, queryParams map[string]string) (*http.Response, error) {
	var req *http.Request
	var err error
	contentType := "application/json"

	if payload != nil && httpMethod != "GET" {
		payloadBuf := new(bytes.Buffer)
		for _, header := range headers {
			if header.Key == "Content-Type" {
				contentType = header.Value
			}
		}

		if contentType == "application/x-www-form-urlencoded" {
			data := url.Values{}
			for i, v := range payload.(map[string]interface{}) {
				data.Add(i, v.(string))
			}
			payloadBuf = bytes.NewBufferString(data.Encode())
		} else {
			if err := json.NewEncoder(payloadBuf).Encode(payload); err != nil {
				return nil, err
			}
		}

		req, err = http.NewRequest(httpMethod, fullURL, payloadBuf)
	} else {
		req, err = http.NewRequest(httpMethod, fullURL, nil)
	}

	if err != nil {
		return nil, err
	}

	if httpMethod == http.MethodGet {
		q := req.URL.Query()
		for i, v := range queryParams {
			q.Add(i, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Set("Content-Type", contentType)

	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}

	client := &http.Client{
		Timeout: time.Minute * five,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CallAPIMicroSite is a function for api call
func CallAPIMicroSite(httpMethod, fullURL string, headers []CallerHeader, payload interface{}, queryParams map[string]string) (*http.Response, error) {
	var req *http.Request
	var err error

	contentType := "application/json"

	for _, header := range headers {
		if header.Key == "Content-Type" {
			contentType = header.Value
		}
		log.Println("headers", header)
	}

	req, err = http.NewRequest(httpMethod, fullURL, nil) //nolint:noctx

	if queryParams != nil {
		q := req.URL.Query()
		for i, v := range queryParams {
			log.Println("queyr params", i, v)
			q.Add(i, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if err != nil {
		log.Println("error", err)
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}

	client := &http.Client{
		Timeout: time.Minute * five,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("error", err)
		return nil, err
	}

	log.Println("res", res)
	return res, nil
}
