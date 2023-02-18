// Copyright © 2023 Lancelot

package cmd

import (
	"go-chatgpt-cli/config"
	"go-chatgpt-cli/logger"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get openAI key that you set",
	Run: func(cmd *cobra.Command, args []string) {
		c := &config.KeyConfig{}
		err := config.Load(c)
		if err != nil {
			logger.LanLogger.Error(err)
			return
		}
		cmd.Printf("key:  [%s]\ndate: [%s]\n", c.Key, c.Date)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
