package cmd

import (
	"npmrcm/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "npmrcm",
	Short:             "A cli tool to manage npmrc profiles",
	Version:           version.Value,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.AddCommand(setupCmd, listCmd, activeCmd, switchCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}
