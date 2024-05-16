package root

import (
	"github.com/sikalabs/mon/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "mon",
	Short: "mon, " + version.Version,
}
