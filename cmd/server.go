package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/stuxcd/custos-admission-controller/pkg/server"
)

var (
	server_short = "run HTTP server"
	server_long  = "run HTTP server on specified port"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server [port]",
	Short: server_short,
	Long:  server_long,
	RunE: func(cmd *cobra.Command, args []string) error {
		port := os.Getenv("PORT")
		if port == "" {
			if len(args) < 1 {
				return fmt.Errorf("port is requireds")
			}
			port = args[0]
		}

		if _, err := strconv.Atoi(port); err != nil {
			port = args[0]
		}

		stopCh := server.SetupSignalHandler()
		jsonLogger, _ := NewLogger("info", "json")

		server.ListenAndServe(port, time.Minute, jsonLogger, stopCh)
		return nil
	},
}
