package cli

import (
	"vault-go/utils"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "short description for init",
	Long:  "long description for init",
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.BackupFolderSetup()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
