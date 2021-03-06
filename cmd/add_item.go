package cmd

import (
	"bwca/bitwarden"
	"bwca/common"
	"fmt"

	"github.com/spf13/cobra"
)

// addItemCmd represents the item command
var addItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Add a new item to the bitwarden vault",
	Long: `EXAMPLE:

Create a new bitwarden login item and add it to the favorites list.

  > bwca add item --name bwca-birch \
    --message "This goes into the NOTE field" \
    --username "bwca-user" \
    --password "bwca-password" \
    --favorite

--------OUTPUT--------
Item has been created
----------------------

	NOTE: use techinques to hide your password from your shell history file

EXAMPLE:

Create a new bitwarden login item and assign it to a specified folder. (TODO)

  > bwca add item --name bwca-birch \
    --folder-id <folder_id> \
	--username "bwca-user" \
	--password "bwca-password"

--------OUTPUT--------
Item has been created
----------------------
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		message, _ := cmd.Flags().GetString("message")
		favorite, _ := cmd.Flags().GetBool("favorite")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		folderID, _ := cmd.Flags().GetString("folder-id")
		folder, _ := cmd.Flags().GetString("folder")
		if len(folderID) > 0 || len(folder) > 0 {
			if len(folder) > 0 {
				folderID = getFolderID(folder)
			}
			if len(folderID) > 0 {
				addItem(name, message, favorite, username, password, folderID)
			} else {
				common.Logger.Error("The folder name was not found")
			}
		} else {
			addItem(name, message, favorite, username, password, folderID)
		}
	},
}

func addItem(name string, message string, favorite bool, username string, password string, folderID string) {

	newItem := bitwarden.Newlogin{
		Type:     1,
		Name:     name,
		Notes:    message,
		Favorite: favorite,
		FolderID: folderID,
		Login: bitwarden.NewloginLogin{
			Username: username,
			Password: password,
		},
		Reprompt: 0,
	}

	status, err := bwClient.NewItem(newItem)
	if err != nil {
		common.Logger.WithError(err).Error("Error creating login item")
	}

	if status.Success {
		fmt.Println("Item has been created")
	}

}

func init() {
	addCmd.AddCommand(addItemCmd)

	addItemCmd.Flags().StringP("name", "n", "", "Name of the new item, use easy unique names, like a gitlab slug")
	addItemCmd.Flags().StringP("message", "m", "", "Note/Message to add to the item")
	addItemCmd.Flags().Bool("favorite", false, "Add this item to the list of favorites")
	addItemCmd.Flags().StringP("username", "u", "", "Username of the item")
	addItemCmd.Flags().StringP("password", "p", "", "Password of the item")
	addItemCmd.Flags().String("folder-id", "", "Store the item in the specified folder id")
	addItemCmd.Flags().StringP("folder", "f", "", "Store the item in the specified folder name, name must be unique")

	addItemCmd.MarkFlagRequired("name")
	addItemCmd.MarkFlagRequired("username")
	addItemCmd.MarkFlagRequired("password")
}
