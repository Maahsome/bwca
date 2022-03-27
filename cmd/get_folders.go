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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
