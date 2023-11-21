package collections

import "sync"

type Set[T comparable] struct {
	mem map[T]struct{}
	mu  *sync.RWMutex
}

func NewSet[T comparable](elems ...T) Set[T] {
	s := Set[T]{
		mem: make(map[T]struct{}),
		mu:  &sync.RWMutex{},
	}

	for _, elem := range elems {
		s.Add(elem)
	}

	return s
}

func (s *Set[T]) Add(el T) bool {
	if s.Contains(el) {
		return false
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.mem[el] = struct{}{}
	return true
}

func (s *Set[T]) Contains(el T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, contains := s.mem[el]
	return contains
}

func (s *Set[T]) Remove(el T) bool {
	if !s.Contains(el) {
		return false
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.mem, el)
	return true
}

func (s *Set[T]) Equal(s2 *Set[T]) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	s2.mu.RLock()
	defer s2.mu.RUnlock()

	if len(s.mem) != len(s2.mem) {
		return false
	}

	for el := range s.mem {
		if !s2.Contains(el) {
			return false
		}
	}
	return true
}
