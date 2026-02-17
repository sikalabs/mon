package cron_alert

import (
	"fmt"
	"os"

	"github.com/sikalabs/mon/cmd/alpha"
	"github.com/sikalabs/mon/pkg/alert"
	"github.com/sikalabs/mon/pkg/config"
	"github.com/sikalabs/mon/pkg/http_checks"
	"github.com/sikalabs/mon/pkg/notify"
	"github.com/sikalabs/slu/pkg/utils/error_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "cron-alert",
	Short: "Run alerting for cron",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		config := config.LoadConfig()
		hostname, err := os.Hostname()
		error_utils.HandleError(err)

		alerts, err := alert.GetAlert()
		error_utils.HandleError(err)

		body := alert.SprintAlerts(alerts)
		fmt.Print(body)

		body += "\n"
		_, out := http_checks.RunHttpChecks(config)
		body += out
		fmt.Println()

		err = notify.SendEmailNotification(config, hostname, body)
		error_utils.HandleError(err)

		err = notify.SendTelegramNotification(config, hostname, body)
		error_utils.HandleError(err)
	},
}

func init() {
	alpha.Cmd.AddCommand(Cmd)
}
