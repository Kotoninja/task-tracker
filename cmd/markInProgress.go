/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	storage "github.com/Kotoninja/task-tracker/internal"
	task "github.com/Kotoninja/task-tracker/pkg/modules"
	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [id]",
	Short: "Mark a task as in-progress",
	Long: `Set a task as 'in-progress' by providing its ID. 
This status shows the task is currently being worked on.

Example: task-cli mark-in-progress 7`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please specify the ID")
			return
		}

		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		statusStr := string(task.InProgress)
		if err = storage.StorageIO.Update(id, nil, &statusStr); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
