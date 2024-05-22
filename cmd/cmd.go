package cmd

import (
	_ "github.com/sikalabs/mon/cmd/alpha"
	_ "github.com/sikalabs/mon/cmd/alpha/cron_alert"
	_ "github.com/sikalabs/mon/cmd/alpha/get_info"
	"github.com/sikalabs/mon/cmd/root"
	_ "github.com/sikalabs/mon/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
