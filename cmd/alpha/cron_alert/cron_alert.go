package cron_alert

import (
	"fmt"
	"log"
	"os"

	"github.com/sikalabs/mon/cmd/alpha"
	"github.com/sikalabs/mon/pkg/alert"
	"github.com/sikalabs/mon/pkg/config"
	"github.com/sikalabs/mon/pkg/notify"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "cron-alert",
	Short: "Run alerting for cron",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		config := config.LoadConfig()
		hostname, err := os.Hostname()
		handleError(err)

		alerts, err := alert.GetAlert()
		handleError(err)

		body := alert.SprintAlerts(alerts)
		fmt.Print(body)

		err = notify.SendEmailNotification(config, hostname, body)
		handleError(err)
	},
}

func init() {
	alpha.Cmd.AddCommand(Cmd)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
