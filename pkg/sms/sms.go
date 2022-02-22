package sms

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/geomethu/afitech_svc_communication/config"
)

type SMSMessage interface {
	Send(from, to string) (*SendMessageResponse, error)
}

type SMSConfig struct {
	Env      string
	APIKey   string
	Username string
}

type SMSService struct {
	SMSMessage
	SMSConfig
}

func NewSMSService(cfg *config.Config) *SMSService {
	return &SMSService{
		SMSConfig: SMSConfig{Env: cfg.SMS.Env, APIKey: cfg.SMS.APIKey, Username: cfg.SMS.Username},
	}
}

// Send sms - POST
func (service *SMSService) Send(from, to, message string) (*SendMessageResponse, error) {
	values := url.Values{}
	values.Set("username", service.Username)
	values.Set("to", to)
	values.Set("message", message)
	if from != "" {
		values.Set("from", from)
	}

	smsURL := GetSmsURL(service.Env)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	res, err := service.newPostRequest(smsURL, values, headers)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var smsMessageResponse SendMessageResponse
	if err := json.NewDecoder(res.Body).Decode(&smsMessageResponse); err != nil {
		return nil, errors.New("unable to parse sms response " + err.Error())
	}
	return &smsMessageResponse, nil

}

func (service *SMSService) newPostRequest(url string, values url.Values, headers map[string]string) (*http.Response, error) {
	reader := strings.NewReader(values.Encode())
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Length", strconv.Itoa(reader.Len()))
	req.Header.Set("apikey", service.APIKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}
