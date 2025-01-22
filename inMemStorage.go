package main

type InMemoryStorage struct {
	todos []Todo
}

func (s *InMemoryStorage) Save(todos []Todo) error {
	s.todos = todos
	return nil
}

func (s *InMemoryStorage) Load() ([]Todo, error) {
	return s.todos, nil
}