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
	Short: "Get a list of items from the bitwarden vault",
	Long: `EXAMPLE:

Get a list of items

  > bwca get items

--------OUTPUT--------
ID                                  	NAME              	FOLDERID
24c8a7e0-95d4-4abf-bf1b-ae6700fd6ffd	bwca-birch
486ab4e1-faac-4b2e-bd78-ae6700fe93e9	bwca-dugout
81c92719-4b4a-4d03-8eed-ae6700fed68a	bwca-tule
----------------------
`,
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
