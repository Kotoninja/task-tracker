/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	storage "github.com/Kotoninja/task-tracker/internal"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task by its ID",
	Long: `Delete an existing task from your task tracker using its unique ID.

This command permanently removes a task from your task list based on the
provided task ID. The task will be deleted from storage and cannot be recovered.
The command validates that the ID is a valid number and that the task exists.

Examples:
  task-tracker delete 1
  task-tracker delete 42
  task-tracker delete 100

The command accepts exactly one argument - the numeric ID of the task to delete.
If no ID is provided, the command will prompt you to specify one.
If multiple arguments are given, the command will display an error message.
If the ID does not exist in the task list, an appropriate error will be shown.

After successful deletion, the command will confirm that the task has been
removed from your task tracker.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				if err = storage.StorageIO.Delete(uint64(id)); err != nil {
					fmt.Println(err)
				}
			}
		} else if len(args) == 0 {
			fmt.Println("Specify ID")
		} else {
			fmt.Println("There is more than one argument")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
