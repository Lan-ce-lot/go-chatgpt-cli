// Copyright © 2023 Lancelot

package cmd

import (
	"fmt"
	"go-chatgpt-cli/config"
	"go-chatgpt-cli/services"

	"github.com/spf13/cobra"
)

const Banner = `
                        _           _              _             _ _
  __ _  ___         ___| |__   __ _| |_ __ _ _ __ | |_       ___| (_)
 / _` + "`" + ` |/ _ \ _____ / __| '_ \ / _` + "`" + ` | __/ _` + "`" + ` | '_ \| __|____ / __| | |
| (_| | (_) |_____| (__| | | | (_| | || (_| | |_) | ||_____| (__| | |
 \__, |\___/       \___|_| |_|\__,_|\__\__, | .__/ \__|     \___|_|_|
 |___/                                 |___/|_|

`

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run chatbot",
	Run: func(cmd *cobra.Command, args []string) {
		c := &config.KeyConfig{}
		err := config.Load(&c)
		if err != nil {
			return
		}
		services.InitClient(c.Key)
		fmt.Printf("\u001B[33m%s\u001B[0m", Banner)
		fmt.Printf("go-chatgpt-cli is running, use 'exit' or Ctrl-D (i.e. EOF) to exit\n\n")
		for {
			request := services.PrintUserName(services.DefaultUserName)
			if request == "exit" {
				return
			}
			_, err := services.GetResponse(request)
			if err != nil {
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
