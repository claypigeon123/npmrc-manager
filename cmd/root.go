package cmd

import (
	"github.com/spf13/cobra"
	"npmrcm/version"
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
