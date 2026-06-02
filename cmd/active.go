package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"npmrcm/internal"
)

var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "Describe the current active profile",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runActive()
	},
}

func runActive() error {
	fsUtil := internal.FsUtil{}
	profiles, err := fsUtil.GetProfiles()

	if err != nil {
		return err
	}

	for _, profile := range profiles {
		if profile.Active {
			fmt.Println(profile.Name)
			break
		}
	}

	return nil
}
