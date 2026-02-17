package config

const (
	// mon.yaml schema
	MON_NOTIFICATIONS_SEND_OK = "notifications.send_ok"

	MON_NOTIFICATIONS_EMAIL_SMTP_HOST       = "notifications.mail.smtp_host"
	MON_NOTIFICATIONS_EMAIL_SMTP_PORT       = "notifications.mail.smtp_port"
	MON_NOTIFICATIONS_EMAIL_SMTP_USERNAME   = "notifications.mail.smtp_username"
	MON_NOTIFICATIONS_EMAIL_SMTP_PASSWORD   = "notifications.mail.smtp_password"
	MON_NOTIFICATIONS_EMAIL_SMTP_EMAIL_FROM = "notifications.mail.smtp_email_from"
	MON_NOTIFICATIONS_EMAIL_EMAILS_TO       = "notifications.mail.emails_to"

	MON_NOTIFICATIONS_TELEGRAM_TOKEN    = "notifications.telegram.token"
	MON_NOTIFICATIONS_TELEGRAM_CHAT_IDS = "notifications.telegram.chat_ids"

	MON_HTTP_CHECKS = "http_checks"

	MON_LEGACY_LIMITS_ROOT_DISK_FREE_GB      = "legacy_limits.root_disk_free_gb"
	MON_LEGACY_LIMITS_ROOT_DISK_FREE_PERCENT = "legacy_limits.root_disk_free_percent"
)
