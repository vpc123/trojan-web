package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vpc123/trojan-web/trojan"
)

// tlsCmd represents the tls command
var tlsCmd = &cobra.Command{
	Use:   "tls",
	Short: "证书安装",
	Run: func(cmd *cobra.Command, args []string) {
		trojan.InstallTls()
	},
}

func init() {
	rootCmd.AddCommand(tlsCmd)
}
