package sms

import "fmt"

const (
	SUCCESS    = "success"
	FAILED     = "failed"
	SENT       = "sent"
	QUEUED     = "queued"
	Sandbox    = "sandbox"
	Production = "production"
)

// SendMessageResponse is a model
type SendMessageResponse struct {
	SMS SMS2 `json:"SMSMessageData"`
}

// SMS2 is a model
type SMS2 struct {
	Recipients []Recipient `json:"recipients"`
}

// Recipient is a model
type Recipient struct {
	Number    string `json:"number"`
	Cost      string `json:"cost"`
	Status    string `json:"status"`
	MessageID string `json:"messageId"`
}


// GetAPIHost returns either sandbox or prod
func GetAPIHost(env string) string {
	return getHost(env, "api")
}

// GetSmsURL is the sms endpoint
func GetSmsURL(env string) string {
	return GetAPIHost(env) + "/version1/messaging"
}


func getHost(env, service string) string {
	if env != Sandbox {
		return fmt.Sprintf("https://%s.africastalking.com", service)
	}

	return fmt.Sprintf("https://%s.sandbox.africastalking.com", service)
}
