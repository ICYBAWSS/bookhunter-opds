package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/bookstairs/bookhunter/internal/opds"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the OPDS server",
	Long:  "Start an HTTP server that provides OPDS feeds for book search and download.",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := opds.NewServer()
		addr := ":8080" // Default address
		fmt.Println("OPDS server starting...")
		if err := server.Start(addr); err != nil {
			return err
		}
		// Block until interrupted
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		fmt.Println("OPDS server stopped.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
} 