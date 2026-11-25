package cli

import (
	"fmt"
	"vault-go/internal/backup"

	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "short description for backup",
	Long:  "long description for backup",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("expected backup folder, received none")
		}

		bFolder := args[0]
		if err := backup.BackupProcess(bFolder); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
