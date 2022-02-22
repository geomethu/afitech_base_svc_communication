package config

import (
	"os"

	"github.com/geomethu/afitech_base_svc_communication/pkg/log"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type (
	Config struct {
		SMS ShortMessage
		Email
		Server Server
	}

	Server struct {
		AuthSvcServer ServerConfig
		CommSvcServer ServerConfig
	}

	ServerConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		TLS  bool   `mapstructure:"tls"`
	}

	ShortMessage struct {
		SMSConfig
	}

	SMSConfig struct {
		Username string `required:"true"`
		APIKey   string `required:"true"`
		Env      string `required:"true"`
	}

	Email struct {
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := godotenv.Load("../.env")
	if err != nil {
		log.Logger.Fatal("communication service", zap.String("Load env", err.Error()))
		return nil, err
	}

	cfg.SMS.SMSConfig.APIKey = os.Getenv("SMS_API_KEY")
	cfg.SMS.SMSConfig.Username = os.Getenv("SMS_API_USERNAME")
	cfg.SMS.SMSConfig.Env = os.Getenv("SMS_API_ENV")

	cfg.Server.CommSvcServer.Host = os.Getenv("COMM_SVC_HOST")
	cfg.Server.CommSvcServer.Port = os.Getenv("COMM_SVC_PORT")

	return cfg, nil
}
