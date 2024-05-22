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
		config.SMTPHost,
		strconv.Itoa(config.SMTPPort),
		config.SMTPUsername,
		config.SMTPPassword,
		config.SMTPEmailFrom,
		"ondrejsika@ondrejsika.com",
		"[mon] Alert from "+hostname,
		body+GetEmailFooter(),
	)
	return err
}
