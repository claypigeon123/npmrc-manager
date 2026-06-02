package cmd

import (
	"fmt"
	"npmrcm/internal"

	"github.com/spf13/cobra"
)

var Verbose bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured .npmrc profiles",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runList()
	},
}

func init() {
	listCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Prints additional information about profiles")
}

func runList() error {
	fsUtil := internal.FsUtil{}
	profiles, err := fsUtil.GetProfiles()

	if err != nil {
		return err
	}

	if Verbose {
		fmt.Println("Listing .npmrc profiles:")
	}

	for _, profile := range profiles {
		printProfile(profile)
	}

	return nil
}

func printProfile(profile internal.NpmrcProfile) {
	if !Verbose {
		if profile.Active {
			fmt.Print("---> ")
		} else {
			fmt.Print("     ")
		}
		fmt.Println(profile.Name)

		return
	}

	activeVerb := "no"
	if profile.Active {
		activeVerb = "yes"
	}

	fmt.Println()
	fmt.Printf("# profile name:         %s\n", profile.Name)
	fmt.Printf("# persistent path:      %s\n", profile.Path)
	fmt.Printf("# active?               %s\n", activeVerb)
}
