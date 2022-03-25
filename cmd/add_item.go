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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		message, _ := cmd.Flags().GetString("message")
		favorite, _ := cmd.Flags().GetBool("favorite")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		addItem(name, message, favorite, username, password)
	},
}

func addItem(name string, message string, favorite bool, username string, password string) {

	newItem := bitwarden.Newlogin{
		Type:     1,
		Name:     name,
		Notes:    message,
		Favorite: favorite,
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
	addItemCmd.Flags().BoolP("favorite", "f", false, "Add this item to the list of favorites")
	addItemCmd.Flags().StringP("username", "u", "", "Username of the item")
	addItemCmd.Flags().StringP("password", "p", "", "Password of the item")

	addItemCmd.MarkFlagRequired("name")
	addItemCmd.MarkFlagRequired("username")
	addItemCmd.MarkFlagRequired("password")
}
