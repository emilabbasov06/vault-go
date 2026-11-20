package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "vgo",
	Short:   "short description",
	Long:    "long description",
	Version: "1.10.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "vgo running...\n")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error occurred while executing VaultGo: `%s`\n", err)
		os.Exit(1)
	}
}
