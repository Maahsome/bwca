package cmd

import (
	"bwca/bitwarden"
	"bwca/common"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// getItemsCmd represents the items command
var getItemsCmd = &cobra.Command{
	Use:   "items",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getItems()
	},
}

func getItems() {
	items, err := bwClient.GetItems("")
	if err != nil {
		common.Logger.Fatal("Failed to list GetItems")
	}
	fmt.Println(itemsDataToString(items, fmt.Sprintf("%#v", items)))
}

func itemsDataToString(itemData bitwarden.Items, raw string) string {

	switch strings.ToLower(c.OutputFormat) {
	case "raw":
		return raw
	case "json":
		return itemData.ToJSON()
	case "gron":
		return itemData.ToGRON()
	case "yaml":
		return itemData.ToYAML()
	case "text", "table":
		return itemData.ToTEXT(c.NoHeaders)
	default:
		return itemData.ToTEXT(c.NoHeaders)
	}
}

func init() {
	getCmd.AddCommand(getItemsCmd)
}
