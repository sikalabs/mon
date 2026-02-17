package run

import (
	"github.com/sikalabs/mon/cmd/alpha"
	"github.com/sikalabs/mon/pkg/run"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Run alerting for cron",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		run.RunOrDie()
	},
}

func init() {
	alpha.Cmd.AddCommand(Cmd)
}
