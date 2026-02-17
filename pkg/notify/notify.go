package notify

import (
	"strconv"

	"github.com/ondrejsika/gosendmail/lib"
	"github.com/sikalabs/mon/pkg/config"
)

func GetEmailFooter() string {
	return `
--
mon
`
}

func SendEmailNotification(
	config config.Config,
	hostname string,
	body string,
) error {
	err := lib.GoSendMail(
		config.Notifications.Mail.SMTPHost,
		strconv.Itoa(config.Notifications.Mail.SMTPPort),
		config.Notifications.Mail.SMTPUsername,
		config.Notifications.Mail.SMTPPassword,
		config.Notifications.Mail.SMTPEmailFrom,
		"ondrejsika@ondrejsika.com",
		"[mon] Alert from "+hostname,
		body+GetEmailFooter(),
	)
	return err
}
