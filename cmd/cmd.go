package cmd

import (
	"github.com/sikalabs/mon/cmd/root"
	_ "github.com/sikalabs/mon/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
