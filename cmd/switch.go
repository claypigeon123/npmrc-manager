package cmd

import (
	"errors"
	"fmt"
	"npmrcm/internal"

	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:        "switch <target profile>",
	Short:      "Switch to the specified .npmrc profile",
	Args:       cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ArgAliases: []string{"profile"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSwitch(args[0])
	},
}

func runSwitch(targetProfileName string) error {
	fsUtil := internal.FsUtil{}
	profiles, err := fsUtil.GetProfiles()

	if err != nil {
		return err
	}

	var targetProfile *internal.NpmrcProfile

	for _, profile := range profiles {
		if profile.Name != targetProfileName {
			continue
		}

		targetProfile = &profile
		break
	}

	if targetProfile == nil {
		msg := fmt.Sprintf("Target profile [%s] does not exist.\nSUGGESTION: List available profiles with \"npmrcm list\" or \"npmrcm list --verbose\"\n", targetProfileName)
		return errors.New(msg)
	}

	if targetProfile.Active {
		fmt.Printf("Profile [%s] is already active\n", targetProfileName)
		return nil
	}

	err = fsUtil.ReplaceNpmrcContent(targetProfile.Content)

	if err != nil {
		return err
	}

	fmt.Printf("Switched to profile [%s]\n", targetProfileName)
	return nil
}
