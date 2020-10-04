package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func flag(file string) {
	parts := strings.Split(file, "-")
	name := strings.Trim(strings.ReplaceAll(parts[0], "_", " "), " ")

	fmt.Println(fmt.Sprintf("country: %+v", name))

	c, err := country.FindName(translateName(name))
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Flag)
	DownloadFile(filepath.Join("icons", file)+".svg", c.Flag)
}

func translateName(name string) string {
	switch {
	case strings.HasPrefix(name, "Korea"):
		return "Republic of Korea"
	case strings.HasPrefix(name, "United States"):
		return "United States of America"
	}

	return name
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
