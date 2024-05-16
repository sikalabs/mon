package alpha

import (
	"github.com/sikalabs/mon/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "alpha",
	Short:   "alpha features",
	Args:    cobra.NoArgs,
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
