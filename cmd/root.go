// Copyright © 2023 Lancelot

package cmd

import (
	"go-chatgpt-cli/logger"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-chatgpt-cli",
	Short: "A brief description of your application",
	Long: `
Go-chatgpt-cli is a CLI to communication to gpt3.

`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	cmd.Execute()
	//
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if e := rootCmd.Execute(); e != nil {
		logger.LanLogger.Error(e)
		os.Exit(1)
	}
}

func init() {
}
