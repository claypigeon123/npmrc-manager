package cmd

import (
	"fmt"
	"npmrcm/internal"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Run guided setup of required configuration for this tool",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSetup()
	},
}

func runSetup() error {
	fsUtil := internal.FsUtil{}
	fmt.Println("Initializing guided setup")

	configured := fsUtil.IsConfigured()

	if configured {
		configDir, err := fsUtil.GetAppConfigDir()

		if err != nil {
			return err
		}

		fmt.Printf("Npmrcm is already configured at path %s\n", configDir)
		return nil
	}

	return nil
}
