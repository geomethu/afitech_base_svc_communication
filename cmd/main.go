package main

import (
	"github.com/geomethu/afitech_svc_communication/config"
	"github.com/geomethu/afitech_svc_communication/internal/app"
	"github.com/geomethu/afitech_svc_communication/pkg/log"
	"github.com/geomethu/afitech_svc_communication/pkg/sms"
	"go.uber.org/zap"
)

func main() {
	log.InitLogger()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Logger.Fatal("communication service", zap.Error(err))
	}

	sms := sms.NewSMSService(cfg)
	_ = sms

	app.Run(cfg)

}
