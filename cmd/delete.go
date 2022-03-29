package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"remove"},
	Short:   "Delete objects from the bitwarden vault",
	Long: `Delete various objects from the bitwarden vault:
  - item
  - folder
`,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
