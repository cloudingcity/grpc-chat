package cmd

import (
	"github.com/cloudingcity/grpc-chat/internal/client"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run as client",
	Run: func(cmd *cobra.Command, args []string) {
		client.Connect(addr, username, password)
	},
}

var (
	username string
	addr     string
)

func init() {
	clientCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	clientCmd.Flags().SortFlags = false
	clientCmd.Flags().StringVarP(&username, "username", "u", "", "the client username")
	clientCmd.Flags().StringVarP(&password, "password", "p", "", "the client password")
	clientCmd.Flags().StringVarP(&addr, "addr", "a", ":8888", "the server address")
	_ = clientCmd.MarkFlagRequired("username")
	_ = clientCmd.MarkFlagRequired("password")
}
