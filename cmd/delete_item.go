package cmd

import (
	"bwca/common"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteItemCmd represents the item command
var deleteItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Delete an item from the bitwarden vault",
	Long: `EXAMPLE:

Delete a bitwarden login item from the bitwarden vault.

  > bwca delete item --item-id <item_id>

--------OUTPUT--------
Item <item_id> has been deleted
----------------------
`,
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
