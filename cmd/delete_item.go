package cmd

import (
	"bwca/common"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteItemCmd represents the item command
var deleteItemCmd = &cobra.Command{
	Use:   "item",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		itemID, _ := cmd.Flags().GetString("item-id")
		deleteItem(itemID)
	},
}

func deleteItem(id string) {

	ok, err := bwClient.DeleteItem(id)
	if err != nil {
		common.Logger.WithError(err).Error(fmt.Sprintf("Failed to delete item %s", id))
	}
	if ok {
		fmt.Printf("Item %s has been deleted\n", id)
	} else {
		fmt.Printf("Failed to delete item %s\n", id)
	}

}

func init() {
	deleteCmd.AddCommand(deleteItemCmd)

	deleteItemCmd.Flags().StringP("item-id", "i", "", "The ID of the item to delete")
	deleteItemCmd.MarkFlagRequired("item-id")
}
