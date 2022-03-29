package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create", "new"},
	Short:   "Create new objects in the bitwarden vault",
	Long: `Add various object to the bitwarden vault:
  - folder
  - item
.`,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
