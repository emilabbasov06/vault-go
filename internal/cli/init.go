package cli

import (
	"vault-go/internal/config"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "short description for init",
	Long:  "long description for init",
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := config.FolderSetup()
		if err != nil {
			return err
		}

		rootCmd.Version = repo.Version
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
