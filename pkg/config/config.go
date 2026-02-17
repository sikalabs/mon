package config

import (
	"strings"

	"github.com/spf13/viper"
)

type HTTPCheck struct {
	URL string `mapstructure:"url"`
}

type Config struct {
	Meta struct {
		Version int `mapstructure:"version"`
	} `mapstructure:"meta"`
	Notifications struct {
		SendOK bool `mapstructure:"send_ok"`
		Mail   struct {
			SMTPHost      string   `mapstructure:"smtp_host"`
			SMTPPort      int      `mapstructure:"smtp_port"`
			SMTPUsername  string   `mapstructure:"smtp_username"`
			SMTPPassword  string   `mapstructure:"smtp_password"`
			SMTPEmailFrom string   `mapstructure:"smtp_email_from"`
			EmailsTo      []string `mapstructure:"emails_to"`
		} `mapstructure:"mail"`
		Telegram struct {
			Token   string  `mapstructure:"token"`
			ChatIDs []int64 `mapstructure:"chat_ids"`
		} `mapstructure:"telegram"`
	} `mapstructure:"notifications"`
	HTTPChecks   []HTTPCheck `mapstructure:"http_checks"`
	LegacyLimits struct {
		RootDiskFreeGB      int `mapstructure:"root_disk_free_gb"`
		RootDiskFreePercent int `mapstructure:"root_disk_free_percent"`
	} `mapstructure:"legacy_limits"`
}

func LoadConfig() Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("mon")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.SetEnvPrefix("mon")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("legacy_limits.root_disk_free_gb", 5)
	viper.SetDefault("legacy_limits.root_disk_free_percent", 10)

	var c Config
	viper.Unmarshal(&c)

	return c
}
