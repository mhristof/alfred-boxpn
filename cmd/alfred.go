package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/mhristof/go-alfred"
	"github.com/spf13/cobra"
)

var (
	alfredCmd = &cobra.Command{
		Use:   "alfred",
		Short: "List alfred options",
		Run: func(cmd *cobra.Command, args []string) {
			configs := "boxpn-openvpn-configs"
			files, err := ioutil.ReadDir(configs)
			if err != nil {
				panic(err)
			}

			var opts alfred.ScriptFilter
			for _, file := range files {

				item := opts.Add(file.Name(), file.Name())
				abs, err := filepath.Abs(filepath.Join(configs, file.Name()))
				if err != nil {
					panic(err)
				}

				item.SetArg(fmt.Sprintf("openvpn '%s'", strings.ReplaceAll(abs, " ", "\\ ")))
			}

			creds := opts.Add("creds", "setup creds, required args are 'username password'")
			creds.SetArg("creds")

			creds = opts.Add("close", "Close the running connection")
			creds.SetArg("close")

			opts.Print()
		},
	}
)

func init() {
	rootCmd.AddCommand(alfredCmd)
}
