package cmd

import (
	"bwca/bitwarden"
	"bwca/common"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// getFolderCmd represents the folder command
var getFolderCmd = &cobra.Command{
	Use:   "folder",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		folderID, _ := cmd.Flags().GetString("folder-id")
		folderName, _ := cmd.Flags().GetString("folder-name")

		if !c.FormatOverridden {
			c.OutputFormat = "json"
		}
		if len(folderName) > 0 {
			folderID = getFolderID(folderName)
		}
		if len(folderID) > 0 {
			getFolder(folderID)
		} else {
			common.Logger.Error("You must specify --item-id or --item-name")
		}
	},
}

func getFolderID(name string) string {

	id, err := bwClient.FindFolder(name)
	if err != nil {
		common.Logger.Error("Could not find the item")
	}
	return id
}

func getFolder(id string) {
	folder, err := bwClient.GetFolder(id)
	if err != nil {
		common.Logger.Fatal("Failed to list GetItem")
	}
	fmt.Println(folderDataToString(folder, fmt.Sprintf("%#v", folder)))
}

func folderDataToString(folderData bitwarden.Folder, raw string) string {

	switch strings.ToLower(c.OutputFormat) {
	case "raw":
		return raw
	case "json":
		return folderData.ToJSON()
	case "gron":
		return folderData.ToGRON()
	case "yaml":
		return folderData.ToYAML()
	case "text", "table":
		return folderData.ToTEXT(c.NoHeaders)
	default:
		return folderData.ToTEXT(c.NoHeaders)
	}
}

func init() {
	getCmd.AddCommand(getFolderCmd)
	getFolderCmd.Flags().StringP("folder-id", "i", "", "The ID of the item to fetch")
	getFolderCmd.Flags().StringP("folder-name", "n", "", "The name of the item to fetch, careful: name items wisely")
}
