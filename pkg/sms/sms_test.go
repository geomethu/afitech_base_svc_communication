package sms

import (
	"fmt"
	"testing"

	"github.com/geomethu/afitech_svc_communication/config"
	"github.com/stretchr/testify/assert"
)

func TestSmsSend(t *testing.T) {
	sc := config.SMSConfig{
		Env:    "sandbox",
		APIKey: "",
		Username: "sandbox",
	}
	sm:=config.ShortMessage{
		SMSConfig: sc,
	}
	c := config.Config{
		SMS: sm,
	}
	sms_svc := NewSMSService(&c)

	res, err := sms_svc.Send(sc.Username, "+254724248784", "test message")
	fmt.Println(res)
	assert.Nil(t, err)
}
