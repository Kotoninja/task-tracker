package main

import (
	"fmt"

	"github.com/Kotoninja/task-tracker/cmd"
	"github.com/Kotoninja/task-tracker/internal"
)

func main() {
	store, err := storage.NewStore("storage.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	storage.StorageIO = store
	cmd.Execute()
}
