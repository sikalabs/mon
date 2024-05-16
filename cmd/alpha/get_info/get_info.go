package get_info

import (
	"log"

	"github.com/sikalabs/mon/cmd/alpha"
	"github.com/sikalabs/mon/pkg/collect"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "get-info",
	Short:   "Get Info",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := collect.PrintHostInfo()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	alpha.Cmd.AddCommand(Cmd)
}
