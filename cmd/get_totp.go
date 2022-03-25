package cmd

import (
	"bwca/common"
	"fmt"

	"github.com/spf13/cobra"
)

// getTotpCmd represents the password command
var getTotpCmd = &cobra.Command{
	Use:   "totp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
