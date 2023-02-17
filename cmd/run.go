/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-chatgpt-cli/config"
	"go-chatgpt-cli/services"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run chatbot",
	Run: func(cmd *cobra.Command, args []string) {
		c := &config.KeyConfig{}
		config.Load(&c)
		services.InitClient(c.Key)
		fmt.Println("go-chatgpt-cli is running, use 'exit' or Ctrl-D (i.e. EOF) to exit")
		for {
			request := services.PrintUserName(services.DefaultUserName)
			if request == "exit" {
				return
			}
			services.GetResponse(request)
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
