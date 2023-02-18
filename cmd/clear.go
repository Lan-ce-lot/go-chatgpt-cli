// Copyright © 2023 Lancelot

package cmd

import (
	"go-chatgpt-cli/config"
	"go-chatgpt-cli/logger"

	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear openAI key",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Cleanup()
		if err != nil {
			logger.LanLogger.Error(err)
			return
		}
		cmd.Printf("clear config in %s\n", config.Path)
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
