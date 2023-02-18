// Copyright © 2023 Lancelot

package cmd

import (
	"go-chatgpt-cli/config"
	"time"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set openAI key",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Save(config.KeyConfig{
			Key:  args[0],
			Date: time.Now(),
		})
		if err != nil {
			return
		}
		cmd.Printf("set key [%s] success, config in %s\n", args[0], config.Path)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
