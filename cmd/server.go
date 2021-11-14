package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/stuxcd/admission-controller/pkg/server"
)

var (
	server_short = "run HTTP server"
	server_long  = "run HTTP server on specified port"
)

var (
	// flags
	port string
)

func init() {
	serverCmd.PersistentFlags().StringVarP(&port, "port", "p", "8443", "port to listen on")

	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: server_short,
	Long:  server_long,
	RunE: func(cmd *cobra.Command, args []string) error {
		portEnv := os.Getenv("PORT")
		if portEnv != "" {
			port = portEnv
		}

		if _, err := strconv.Atoi(port); err != nil {
			return fmt.Errorf("cannot convert %s to int", port)
		}

		stopCh := server.SetupSignalHandler()
		jsonLogger, _ := NewLogger("info", "json")

		server.ListenAndServe(port, time.Minute, jsonLogger, stopCh)
		return nil
	},
}
