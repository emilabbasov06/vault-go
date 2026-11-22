package cli

import (
	"fmt"

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
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
