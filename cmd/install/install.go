package install

import (
	"github.com/sikalabs/mon/cmd/root"
	"github.com/sikalabs/mon/pkg/install"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "install",
	Short: "Install mon as a systemd service",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		install.Install()
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
