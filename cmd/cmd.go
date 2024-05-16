package cmd

import (
	"github.com/sikalabs/mon/cmd/root"
	_ "github.com/sikalabs/mon/cmd/version"
	_ "github.com/sikalabs/mon/cmd/alpha"
	_ "github.com/sikalabs/mon/cmd/alpha/get_info"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
