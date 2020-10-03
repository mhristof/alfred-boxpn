package cmd

import (
	"github.com/mhristof/alfred-boxpn/config"
	"github.com/mhristof/alfred-boxpn/security"
	"github.com/spf13/cobra"
)

var (
	credsCmd = &cobra.Command{
		Use:   "creds",
		Short: "Set the credentials",
		Args:  cobra.ExactValidArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			creds := security.Creds{
				Service: config.Service,
				Account: config.Account,
				Label:   config.Label,
			}

			creds.Set(args[0], args[1])
		},
	}
)

func init() {
	rootCmd.AddCommand(credsCmd)
}
