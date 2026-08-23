/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	storage "github.com/Kotoninja/task-tracker/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task to your task list",
	Long: `Add a new task with the provided description to your task tracker.

This command creates a new task entry in your task list with a unique ID,
the provided description, and sets the status to "pending" by default.
The task is automatically assigned a creation timestamp and saved to storage.

Examples:
  task-tracker add "Buy groceries"
  task-tracker add "Complete project documentation"
  task-tracker add "Schedule team meeting for Friday"

The command accepts exactly one argument - the task description.
If no description is provided or multiple arguments are given,
the command will display an appropriate error message.

After successful addition, the command will display a confirmation
message with the newly created task ID and description.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			output, err := storage.StorageIO.Add(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(output)
			}
		} else if len(args) == 0 {
			fmt.Println("Nothing add")
		} else {
			fmt.Println("There is more than one argument")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
