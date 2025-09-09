package engine

import "errors"

type Stack[T any] struct {
	items []*T
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item *T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}
