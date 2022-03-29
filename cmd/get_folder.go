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
	Short: "Get a folder from the bitwarden vault",
	Long: `EXAMPLE:

Get a single folder from the bitwarden vault by name

  > bwca get folder --folder-name bwca-folder

--------OUTPUT--------
{
	"data": {
	  "id": "8628259a-2323-4454-9826-ae6700f42a62",
	  "name": "bwca-folder",
	  "object": "folder"
	},
	"success": true
  }
----------------------

EXAMPLE:

Get a single folder from the bitwarden vault by ID

  > bwca get folder --folder-id 8628259a-2323-4454-9826-ae6700f42a62

--------OUTPUT--------
{
	"data": {
	  "id": "8628259a-2323-4454-9826-ae6700f42a62",
	  "name": "bwca-folder",
	  "object": "folder"
	},
	"success": true
  }
----------------------
`,
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
