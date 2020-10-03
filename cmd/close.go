package cmd

import (
	"github.com/mhristof/alfred-boxpn/sudo"
	"github.com/spf13/cobra"
)

var (
	closeCmd = &cobra.Command{
		Use:   "close",
		Short: "Close all openvpn connections",
		Run: func(cmd *cobra.Command, args []string) {
			// do something

			sudo.Run("pkill -9 openvpn")
		},
	}
)

func init() {
	rootCmd.AddCommand(closeCmd)
}
