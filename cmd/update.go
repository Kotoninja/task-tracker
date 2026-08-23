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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id] [newDescription]",
	Short: "Update the description of an existing task",
	Long: `Update the description of a task identified by its ID. 
The new description will replace the existing one.

Examples:
  task-cli update 1 "New task description"
  task-cli update 5 "Fix critical bug in payment module"

Arguments:
  [id]              - The unique identifier of the task to update
  [newDescription]  - The new description text (use quotes for spaces)
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Please specify the ID and new description.")
			return
		}

		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		newDescription := &args[1]

		if err = storage.StorageIO.Update(id, newDescription, nil); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
