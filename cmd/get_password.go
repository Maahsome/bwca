package cmd

import (
	"bwca/common"
	"fmt"

	"github.com/spf13/cobra"
)

// getPasswordCmd represents the password command
var getPasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Get a login password from the bitwarden vault",
	Long: `EXAMPLE:

Fetch just the password property from a login object in the bitwarden vault

  > bwca get password --item-name bwca-birch

--------OUTPUT--------
bwca-password
----------------------
`,
	Run: func(cmd *cobra.Command, args []string) {
		itemID, _ := cmd.Flags().GetString("item-id")
		itemName, _ := cmd.Flags().GetString("item-name")
		if len(itemName) > 0 {
			itemID = getItemID(itemName)
		}
		if len(itemID) > 0 {
			getPassword(itemID)
		} else {
			common.Logger.Error("You must specify --item-id or --item-name")
		}
	},
}

func getPassword(id string) {
	password, err := bwClient.GetPassword(id)
	if err != nil {
		common.Logger.Fatal("Failed to GetPassword")
	}
	fmt.Println(password)
}
func init() {
	getCmd.AddCommand(getPasswordCmd)

	getPasswordCmd.Flags().StringP("item-id", "i", "", "The ID of the item to fetch")
	getPasswordCmd.Flags().StringP("item-name", "n", "", "The name of the item to fetch, careful: name items wisely")
}
