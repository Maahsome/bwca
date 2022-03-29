package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"list"},
	Short:   "Get different objects from the bitwarden vault",
	Long: `Get various objects from the bitwarden vault:
  - item(s)
  - folder(s)
  - username
  - password
  - totp
`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
