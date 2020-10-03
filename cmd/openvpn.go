package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/mhristof/alfred-boxpn/config"
	"github.com/mhristof/alfred-boxpn/security"
	"github.com/mhristof/alfred-boxpn/sudo"
	"github.com/spf13/cobra"
)

var (
	openvpnCmd = &cobra.Command{
		Use:   "openvpn",
		Short: "Start openvpn in the background",
		Args:  cobra.ExactArgs(1),
		Run: func(comm *cobra.Command, args []string) {
			// do something

			if _, err := os.Stat(args[0]); os.IsNotExist(err) {
				panic(fmt.Sprintf("Error, file %s does not exist", args[0]))
			}

			creds := security.Creds{
				Service: config.Service,
				Account: config.Account,
				Label:   config.Label,
			}

			err := ioutil.WriteFile(config.Auth, []byte(creds.Get()), 0644)
			if err != nil {
				panic(err)
			}
			path, err := exec.LookPath("openvpn")
			if err != nil {
				panic(err)
			}

			sudo.Run(fmt.Sprintf(`bash -c "%s '%s'; rm auth.txt"`, path, args[0]))
		},
	}
)

func init() {
	rootCmd.AddCommand(openvpnCmd)
}
