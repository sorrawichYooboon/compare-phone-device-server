/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flowech-device-server/server"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start the server",
	Aliases: []string{"s", "sv"},
	RunE:    runServerCmd,
}

func runServerCmd(cmd *cobra.Command, args []string) error {
	return server.Start()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
