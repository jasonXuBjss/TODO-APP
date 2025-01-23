package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Storage interface {
	Save(todos []Todo) error
	Load() ([]Todo, error)
}

type JSONStorage struct {
	FileName string
	mu sync.Mutex
}

func (s *JSONStorage) Save(todos []Todo) error {
	s.mu.Lock()         
	defer s.mu.Unlock()

	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal todos: %s", err)
	}

	err = os.WriteFile(s.FileName, fileData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %s", err)
	}

	return nil
}

func (s *JSONStorage) Load() ([]Todo, error) {
	s.mu.Lock()         
	defer s.mu.Unlock()
	
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, fmt.Errorf("failed to read file: %s", err)
	}

	var todos []Todo
	err = json.Unmarshal(fileData, &todos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal todos: %s", err)
	}

	return todos, nil
}
