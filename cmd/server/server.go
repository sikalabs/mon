package server

import (
	"time"

	"github.com/sikalabs/mon/cmd/root"
	"github.com/sikalabs/mon/pkg/server"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Run mon server (runs every 10 minutes)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		server.Server(server.ServerOpts{
			WaitTime: 10 * time.Minute,
		})
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
