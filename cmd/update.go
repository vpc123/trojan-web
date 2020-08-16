package cmd

import (
	"github.com/spf13/cobra"
	"trojan/trojan"
)

// upgradeCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新trojan",
	Run: func(cmd *cobra.Command, args []string) {
		trojan.InstallTrojan()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
