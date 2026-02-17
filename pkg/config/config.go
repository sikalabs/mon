package config

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type HTTPCheck struct {
	URL string `mapstructure:"url"`
}

type Config struct {
	Meta struct {
		Version int
	}
	Notifications struct {
		SendOK bool
		Mail   struct {
			SMTPHost      string
			SMTPPort      int
			SMTPUsername  string
			SMTPPassword  string
			SMTPEmailFrom string
			EmailsTo      []string
		}
		Telegram struct {
			Token   string
			ChatIDs []int64
		}
	}
	HTTPChecks []HTTPCheck
}

func LoadConfig() Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("mon")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.SetEnvPrefix("mon")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// notifications
	viper.BindEnv(MON_NOTIFICATIONS_SEND_OK)

	// email
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_HOST)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_PORT)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_USERNAME)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_PASSWORD)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_EMAIL_FROM)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_EMAILS_TO)

	// telegram
	viper.BindEnv(MON_NOTIFICATIONS_TELEGRAM_TOKEN)
	viper.BindEnv(MON_NOTIFICATIONS_TELEGRAM_CHAT_IDS)

	var c Config

	// notifications
	c.Notifications.SendOK = viper.GetBool(MON_NOTIFICATIONS_SEND_OK)

	// email
	c.Notifications.Mail.SMTPHost = viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_HOST)
	c.Notifications.Mail.SMTPPort = viper.GetInt(MON_NOTIFICATIONS_EMAIL_SMTP_PORT)
	c.Notifications.Mail.SMTPUsername = viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_USERNAME)
	c.Notifications.Mail.SMTPPassword = viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_PASSWORD)
	c.Notifications.Mail.SMTPEmailFrom = viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_EMAIL_FROM)
	c.Notifications.Mail.EmailsTo = viper.GetStringSlice(MON_NOTIFICATIONS_EMAIL_EMAILS_TO)

	// telegram
	c.Notifications.Telegram.Token = viper.GetString(MON_NOTIFICATIONS_TELEGRAM_TOKEN)
	for _, s := range viper.GetStringSlice(MON_NOTIFICATIONS_TELEGRAM_CHAT_IDS) {
		id, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			c.Notifications.Telegram.ChatIDs = append(c.Notifications.Telegram.ChatIDs, id)
		}
	}

	// http_checks
	viper.UnmarshalKey(MON_HTTP_CHECKS, &c.HTTPChecks)

	return c
}
