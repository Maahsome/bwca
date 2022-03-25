package cmd

import (
	"bwca/bitwarden"
	"bwca/common"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// getItemCmd represents the item command
var getItemCmd = &cobra.Command{
	Use:   "item",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		itemID, _ := cmd.Flags().GetString("item-id")
		itemName, _ := cmd.Flags().GetString("item-name")

		if !c.FormatOverridden {
			c.OutputFormat = "json"
		}
		if len(itemName) > 0 {
			itemID = getItemID(itemName)
		}
		if len(itemID) > 0 {
			getItem(itemID)
		} else {
			common.Logger.Error("You must specify --item-id or --item-name")
		}
	},
}

func getItemID(name string) string {

	id, err := bwClient.FindItem(name)
	if err != nil {
		common.Logger.Error("Could not find the item")
	}
	return id
}

func getItem(id string) {
	item, err := bwClient.GetItem(id)
	if err != nil {
		common.Logger.Fatal("Failed to list GetItem")
	}
	fmt.Println(itemDataToString(item, fmt.Sprintf("%#v", item)))
}

func itemDataToString(itemData bitwarden.Item, raw string) string {

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
	getCmd.AddCommand(getItemCmd)
	getItemCmd.Flags().StringP("item-id", "i", "", "The ID of the item to fetch")
	getItemCmd.Flags().StringP("item-name", "n", "", "The name of the item to fetch, careful: name items wisely")
}
