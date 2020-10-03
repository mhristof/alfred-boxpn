package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mhristof/alfred-boxpn/config"
	"github.com/spf13/cobra"
)

var (
	fixCmd = &cobra.Command{
		Use:   "fix",
		Short: "Fix credentials to include auth.txt file",
		Run: func(cmd *cobra.Command, args []string) {
			files, err := ioutil.ReadDir(config.Configs)
			if err != nil {
				panic(err)
			}

			for _, file := range files {
				path := filepath.Join(config.Configs, file.Name())
				f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Println(err)
				}
				defer f.Close()
				if _, err := f.WriteString("\nauth-user-pass " + config.Auth); err != nil {
					log.Println(err)
				}

				err = os.Rename(path, strings.ReplaceAll(path, " ", "_"))
				if err != nil {
					panic(err)
				}

			}
		},
	}
)

func init() {
	rootCmd.AddCommand(fixCmd)
}
