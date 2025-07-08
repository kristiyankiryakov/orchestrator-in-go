/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(managerCmd)

	managerCmd.Flags().StringP("host", "H", "0.0.0.0", "Hostname or IP address")
	managerCmd.Flags().IntP("port", "p", 5555, "Port on which to listen")
	managerCmd.Flags().StringSliceP("workers", "w", []string{"localhost:5556"}, "List of workers on which the manager will schedule tasks.")
	managerCmd.Flags().StringP("scheduler", "s", "epvm", "Name of scheduler to use.")
	managerCmd.Flags().StringP("dbType", "d", "memory", "Type of datastore to use for events and tasks (\"memory\" or \"persistent\")")
}

// managerCmd represents the manager command
var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "Manager command to operate a Cube manager node.",
	Long: `cube manager command.

The manager controls the orchestration system and is responsible for:
- Accepting tasks from users
- Scheduling tasks onto worker nodes
- Rescheduling tasks in the event of a node failure
- Periodically polling workers to get task updates`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("manager called")
	},
}
