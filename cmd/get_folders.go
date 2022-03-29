package cmd

import (
	"bwca/bitwarden"
	"bwca/common"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// getFoldersCmd represents the folders command
var getFoldersCmd = &cobra.Command{
	Use:   "folders",
	Short: "Get a list of folders from the bitwarden vault",
	Long: `EXAMPLE:

Get a list of folders.

  > bwca get folders

--------OUTPUT--------
ID                                  	NAME
8628259a-2323-4454-9826-ae6700f42a62	bwca-folder
ebd916e7-cfad-46d1-94da-ae6600322aa2	bwca-paddle
9b32e7b4-1450-476f-9b10-ae660021e5c4	bwca-portage
----------------------

Get a list of folders and suppress the header row

  > bwca get folders --no-headers

--------OUTPUT--------
8628259a-2323-4454-9826-ae6700f42a62	bwca-folder
ebd916e7-cfad-46d1-94da-ae6600322aa2	bwca-paddle
9b32e7b4-1450-476f-9b10-ae660021e5c4	bwca-portage
----------------------
`,
	Run: func(cmd *cobra.Command, args []string) {
		getFolders()
	},
}

func getFolders() {
	folders, err := bwClient.GetFolders()
	if err != nil {
		common.Logger.Fatal("Failed to list GetItems")
	}
	fmt.Println(foldersDataToString(folders, fmt.Sprintf("%#v", folders)))
}

func foldersDataToString(folderData bitwarden.Folders, raw string) string {

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
	getCmd.AddCommand(getFoldersCmd)

}
