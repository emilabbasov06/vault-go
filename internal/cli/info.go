package cli

import (
	"fmt"
	"os"
	"vault-go/internal/config"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "short description for info",
	Long:  "long description for info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, config.GetRepoInfo())
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
