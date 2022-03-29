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
	Short: "Add a new folder to the bitwarden vault",
	Long: `EXAMPLE:
  > bwca add folder --name bwca-new-folder

Folder has been created
`,
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
		fmt.Println("Folder has been created")
	}

}

func init() {
	addCmd.AddCommand(addFolderCmd)
	addFolderCmd.Flags().StringP("name", "n", "", "Name of the new folder, use easy unique names, like a gitlab slug")
	addFolderCmd.MarkFlagRequired("name")
}
