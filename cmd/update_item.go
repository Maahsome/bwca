package cmd

import (
	"bwca/common"
	"fmt"

	"github.com/spf13/cobra"
)

// updateItemCmd represents the item command
var updateItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Update the properties of a bitwarden login item",
	Long: `EXAMPLE:

Update the username of a login object

  > bwca update item --item-name bwca-birch --username birch

EXAMPLE:

Update the password of a login object

  > bwca update item --item-name bwca-birch --password "....bark...."

EXAMPLE:

Update the TOTP KEY of a login object

  > bwca update item --item-name bwca-birch --totp "JBSWY3DPEHPK3PXP"

EXAMPLE:

Update all of the fields, by item ID

  > bwca update item --item-id 24c8a7e0-95d4-4abf-bf1b-ae6700fd6ffd \
      --username birch \
	  --password "....bark...." \
	  --totp "JBSWY3DPEHPK3PXP" \
	  --favorite \
	  --message "Appended notes"

EXAMPLE:

Replace the nodes content

  > bwca upate-item --item-name bwca-birch \
      --message "Replace notes" \
	  --replace-message
`,
	Run: func(cmd *cobra.Command, args []string) {
		itemID, _ := cmd.Flags().GetString("item-id")
		itemName, _ := cmd.Flags().GetString("item-name")
		message, _ := cmd.Flags().GetString("message")
		favorite, _ := cmd.Flags().GetBool("favorite")
		replaceMessage, _ := cmd.Flags().GetBool("replace-message")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		folderID, _ := cmd.Flags().GetString("folder-id")
		folder, _ := cmd.Flags().GetString("folder")
		if len(folderID) > 0 || len(folder) > 0 {
			if len(folder) > 0 {
				folderID = getFolderID(folder)
			}
			if len(folderID) == 0 {
				common.Logger.Fatal("The folder name was not found, cannot proceed")
			}
		}
		if len(itemName) > 0 {
			itemID = getItemID(itemName)
		}
		if len(itemID) > 0 {
			updateItem(itemID, message, favorite, replaceMessage, username, password, folderID)
		} else {
			common.Logger.Error("You must specify --item-id or --item-name")
		}
	},
}

func updateItem(id string, message string, favorite bool, replaceMessage bool, username string, password string, folderID string) {

	item, err := bwClient.GetItem(id)
	if err != nil {
		common.Logger.Fatal("Failed to list GetItem")
	}

	if len(message) > 0 {
		if replaceMessage {
			item.Data.Notes = message
		} else {
			if len(item.Data.Notes) > 0 {
				item.Data.Notes = fmt.Sprintf("%s\n%s", item.Data.Notes, message)
			} else {
				item.Data.Notes = message
			}
		}
	}
	if len(username) > 0 {
		item.Data.Login.Username = username
	}
	if len(password) > 0 {
		item.Data.Login.Password = password
	}
	if len(folderID) > 0 {
		item.Data.FolderID = folderID
	}
	if favorite {
		item.Data.Favorite = favorite
	}

	updatedItem := item.Data

	status, err := bwClient.UpdateItem(id, updatedItem)
	if err != nil {
		common.Logger.WithError(err).Error("Error updating the login item")
	}

	if status.Success {
		fmt.Println("Item has been updated")
	} else {
		fmt.Println("Failed to update the item")
	}
}

func init() {
	updateCmd.AddCommand(updateItemCmd)
	updateItemCmd.Flags().StringP("item-id", "i", "", "Object ID of the login item to update")
	updateItemCmd.Flags().StringP("item-name", "n", "", "Name of the login item to update")
	updateItemCmd.Flags().StringP("message", "m", "", "Note/Message to append/replace")
	updateItemCmd.Flags().Bool("favorite", false, "Add this item to the list of favorites")
	updateItemCmd.Flags().BoolP("replace-message", "r", false, "Add this item to the list of favorites")
	updateItemCmd.Flags().StringP("username", "u", "", "Username of the item")
	updateItemCmd.Flags().StringP("password", "p", "", "Password of the item")
	updateItemCmd.Flags().String("folder-id", "", "Store the item in the specified folder id")
	updateItemCmd.Flags().StringP("folder", "f", "", "Store the item in the specified folder name, name must be unique")
}
