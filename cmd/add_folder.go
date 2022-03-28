package cmd

import (
	"bwca/bitwarden"
	"bwca/common"
	"fmt"

	"github.com/spf13/cobra"
)

// addFolderCmd represents the folder command
var addFolderCmd = &cobra.Command{
	Use:   "folder",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		addFolder(name)
	},
}

func addFolder(name string) {

	newFolder := bitwarden.Newfolder{
		Name: name,
	}

	status, err := bwClient.NewFolder(newFolder)
	if err != nil {
		common.Logger.WithError(err).Error("Error creating login item")
	}

	if status.Success {
		fmt.Println("Item has been created")
	}

}

func init() {
	addCmd.AddCommand(addFolderCmd)
	addFolderCmd.Flags().StringP("name", "n", "", "Name of the new folder, use easy unique names, like a gitlab slug")
	addFolderCmd.MarkFlagRequired("name")
}
