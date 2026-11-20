package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "short description for init",
	Long:  "long description for init",
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't find repo.")
			return err
		}

		fmt.Fprintf(os.Stdout, "Initializing [%s]...\n", repo)
		return nil
	},
}

func init() {
	initCmd.Flags().String("repo", ".", "repo name to initialize")
	rootCmd.AddCommand(initCmd)
}
