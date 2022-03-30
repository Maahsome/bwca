package cmd

import (
	"bwca/common"
	"fmt"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync the local bitwarden vault with the mother ship",
	Long: `EXAMPLE:

Syncronize changes from the cloud vault to the local bitwarden vault

  > bwca sync

--------OUTPUT--------

----------------------
`,
	Run: func(cmd *cobra.Command, args []string) {
		syncVault()
	},
}

func syncVault() {

	status, err := bwClient.Sync()
	if err != nil {
		common.Logger.WithError(err).Error("Error syncing the local bitwarden vault copy")
	}

	if status.Success {
		fmt.Println(status.Data.Title)
	} else {
		fmt.Printf("%s (%s)", status.Data.Title, status.Data.Message)
	}
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
