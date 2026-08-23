package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/Kotoninja/task-tracker/pkg/modules"
)

var (
	StorageIO *Storage
)

type Storage struct {
	filePath string
	pk       uint64
	mx       *sync.Mutex
	tasks    map[uint64]modules.Task
}

func NewStore(filePath string) (*Storage, error) {
	s := &Storage{filePath: filePath, pk: 1, mx: &sync.Mutex{}, tasks: map[uint64]modules.Task{}}
	err := s.load()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return s, nil
}

func (s *Storage) load() error {
	s.mx.Lock()
	defer s.mx.Unlock()

	file, err := os.Open(s.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&s.tasks)
	for _, task := range s.tasks {
		if task.Id >= s.pk {
			s.pk = task.Id + 1
		}
	}
	return nil
}

func (s *Storage) save() error {
	file, err := os.OpenFile(s.filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err = encoder.Encode(s.tasks); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Add(description string) (string, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	newTask := modules.NewTask(s.pk, description)
	s.pk += 1

	s.tasks[newTask.Id] = *newTask
	if err := s.save(); err != nil {
		return "", err
	}
	return fmt.Sprintf("Task added successfully (ID: %d)\n", newTask.Id), nil
}
