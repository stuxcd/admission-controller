package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: version_short,
	Long:  version_long,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("serving http\n")
		return nil
	},
}
