/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	storage "github.com/Kotoninja/task-tracker/internal"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	status = map[string]int{"todo": 0, "in-progress": 1, "done": 2}
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all tasks in your task list",
	Long: `Display all tasks currently stored in your task tracker.

This command retrieves and displays a complete list of all tasks in your
task tracker, showing their ID, description, status, and timestamps.
The tasks are displayed in a formatted table for easy reading.

Examples:
  task-tracker list
  task-tracker list --status done
  task-tracker list --status pending
  task-tracker list --status in-progress

The command supports optional flags to filter tasks by their status.
Without any flags, it displays all tasks regardless of their status.
The output includes task IDs, descriptions, current status, creation time,
and last update time.

If no tasks exist in the tracker, the command will display an appropriate
message indicating that the task list is empty.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			fmt.Println("Too many arguments")
			return
		} else if len(args) == 1 {
			if _, ok := status[args[0]]; ok {
			}
		} else {
			data := [][]string{
				{"Id", "Description", "Status", "Created at", "Updated at"},
			}

			data = append(data, storage.StorageIO.List(nil)...)
			table := tablewriter.NewWriter(os.Stdout)
			table.Header(data[0])
			table.Bulk(data[1:])
			table.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
