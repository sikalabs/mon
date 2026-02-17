package notify

import (
	"log"
	"strconv"

	"github.com/ondrejsika/gosendmail/lib"
	"github.com/sikalabs/mon/pkg/config"
	"github.com/sikalabs/mon/version"
	"github.com/sikalabs/slu/utils/telegram_utils"
)

func GetEmailFooter() string {
	return `
--
mon, ` + version.Version + `
`
}

func SendEmailNotification(
	config config.Config,
	hostname string,
	body string,
) error {
	for _, emailTo := range config.Notifications.Mail.EmailsTo {
		log.Printf("-> Sending email notification to: %s\n", emailTo)
		err := lib.GoSendMail(
			config.Notifications.Mail.SMTPHost,
			strconv.Itoa(config.Notifications.Mail.SMTPPort),
			config.Notifications.Mail.SMTPUsername,
			config.Notifications.Mail.SMTPPassword,
			config.Notifications.Mail.SMTPEmailFrom,
			emailTo,
			"[mon] Alert from "+hostname,
			body+GetEmailFooter(),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func SendTelegramNotification(
	config config.Config,
	hostname string,
	body string,
) error {
	for _, chatID := range config.Notifications.Telegram.ChatIDs {
		log.Printf("-> Sending Telegram notification to: %d\n", chatID)
		err := telegram_utils.TelegramSendMessage(
			config.Notifications.Telegram.Token,
			chatID,
			"[mon] Alert from "+hostname+"\n\n"+body+GetEmailFooter(),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
