package cmd

import (
	"github.com/cloudingcity/grpc-chat/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run as server",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.New(password)
		s.Listen(port)
	},
}

var (
	port     int
	password string
)

func init() {
	serverCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	serverCmd.Flags().SortFlags = false
	serverCmd.Flags().IntVarP(&port, "port", "p", 8888, "the port listen on")
	serverCmd.Flags().StringVarP(&password, "password", "", "", "the server password")
	_ = serverCmd.MarkFlagRequired("password")
}
