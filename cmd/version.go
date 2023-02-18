// Copyright © 2023 Lancelot

package cmd

import (
	"go-chatgpt-cli/version"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version of git is printed on the standard output.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("%s version %s\n", rootCmd.Use, version.Version)
		//fmt.Printf("%s version %s\n", rootCmd.Use, version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
