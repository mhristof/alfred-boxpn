package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
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
				path := filepath.Join(configs, file.Name())
				if err != nil {
					panic(err)
				}

				item.SetArg(fmt.Sprintf("openvpn '%s'", path))
				item.SetMatch(match(file.Name(), "-_."))
				icon := filepath.Join("icons", file.Name()) + ".svg"
				fmt.Println(fmt.Sprintf("icon: %+v", icon))

				if _, err := os.Stat(icon); !os.IsNotExist(err) {
					item.SetIcon(icon)
				}
			}

			creds := opts.Add("creds", "setup creds, required args are 'username password'")
			creds.SetArg("creds")

			creds = opts.Add("close", "Close the running connection")
			creds.SetArg("close")

			opts.Print()
		},
	}
)

func match(haystack, needle string) string {
	for _, char := range needle {
		haystack = strings.ReplaceAll(haystack, string(char), " ")
	}
	return haystack
}

func init() {
	rootCmd.AddCommand(alfredCmd)
}
