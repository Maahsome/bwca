package cmd

import (
	"bwca/common"
	"fmt"

	"github.com/spf13/cobra"
)

// getTotpCmd represents the password command
var getTotpCmd = &cobra.Command{
	Use:   "totp",
	Short: "Get the TOTP from a login item in the bitwarden vault",
	Long: `EXAMPLE:

Get just the calculated TOTP for a login item in the bitwarden vault

  > bwca get totp --item-name bwca-birch

--------OUTPUT--------
714091
----------------------
`,
	Run: func(cmd *cobra.Command, args []string) {
		itemID, _ := cmd.Flags().GetString("item-id")
		itemName, _ := cmd.Flags().GetString("item-name")
		if len(itemName) > 0 {
			itemID = getItemID(itemName)
		}
		if len(itemID) > 0 {
			getTotp(itemID)
		} else {
			common.Logger.Error("You must specify --item-id or --item-name")
		}
	},
}

func getTotp(id string) {
	totp, err := bwClient.GetTOTP(id)
	if err != nil {
		common.Logger.Fatal("Failed to GetTOTP")
	}
	fmt.Println(totp)
}
func init() {
	getCmd.AddCommand(getTotpCmd)

	getTotpCmd.Flags().StringP("item-id", "i", "", "The ID of the item to fetch")
	getTotpCmd.Flags().StringP("item-name", "n", "", "The name of the item to fetch, careful: name items wisely")
}
