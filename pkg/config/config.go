package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Meta struct {
		Version int
	}
	Notifications struct {
		Mail struct {
			SMTPHost      string
			SMTPPort      int
			SMTPUsername  string
			SMTPPassword  string
			SMTPEmailFrom string
		}
	}
}

func LoadConfig() Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("mon")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.SetEnvPrefix("mon")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_HOST)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_PORT)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_USERNAME)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_PASSWORD)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_EMAIL_FROM)

	var c Config
	c.Notifications.Mail.SMTPHost = viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_HOST)
	c.Notifications.Mail.SMTPPort = viper.GetInt(MON_NOTIFICATIONS_EMAIL_SMTP_PORT)
	c.Notifications.Mail.SMTPUsername = viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_USERNAME)
	c.Notifications.Mail.SMTPPassword = viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_PASSWORD)
	c.Notifications.Mail.SMTPEmailFrom = viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_EMAIL_FROM)
	return c
}
