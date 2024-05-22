package cron_alert

import (
	"fmt"

	"github.com/sikalabs/mon/cmd/alpha"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "cron-alert",
	Short: "Run alerting for cron",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("TODO")
	},
}

func init() {
	alpha.Cmd.AddCommand(Cmd)
}
