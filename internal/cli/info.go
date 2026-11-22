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
	RunE: func(cmd *cobra.Command, args []string) error {
		info, err := config.GetConf()
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s", info)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
