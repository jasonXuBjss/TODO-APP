package main

import "sync"

type InMemoryStorage struct {
	todos []Todo
	mu    sync.Mutex
}

func (s *InMemoryStorage) Save(todos []Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.todos = todos
	return nil
}

func (s *InMemoryStorage) Load() ([]Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.todos, nil
}