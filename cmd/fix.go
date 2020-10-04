package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mhristof/alfred-boxpn/config"
	"github.com/mhristof/go-country"
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

				//authUserPass(path)
				flag(file.Name())
				err = os.Rename(path, strings.ReplaceAll(path, " ", "_"))
				if err != nil {
					panic(err)
				}

			}
		},
	}
)

func flag(file string) {
	parts := strings.Split(file, "_")
	name := parts[0]

	fmt.Println(fmt.Sprintf("country: %+v", name))

	c, err := country.FindName(name)
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Flag)
}

func authUserPass(path string) {
	contents, err := ioutil.ReadFile(path)

	if strings.Contains(string(contents), "auth-user-pass") {
		fmt.Println("skipping", path)
		return
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\nauth-user-pass " + config.Auth); err != nil {
		log.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(fixCmd)
}
