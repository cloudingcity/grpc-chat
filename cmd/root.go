package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "grpc-chat",
	Short: "Simple Chat implemented with gRPC",
}

func init() {
	RootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	RootCmd.AddCommand(serverCmd)
	RootCmd.AddCommand(clientCmd)
}
