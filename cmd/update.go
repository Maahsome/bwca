/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update various bitwarded vault objects",
	Long: `Update the following objects:
  - item
`,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
