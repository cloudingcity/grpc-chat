package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "grpc-chat",
	Short: "Simple Chat implemented with gRPC",
}

func init() {
	cobra.OnInitialize(initLog)
	RootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	RootCmd.AddCommand(serverCmd)
	RootCmd.AddCommand(clientCmd)
}

func initLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 03:04:05",
	})
}
