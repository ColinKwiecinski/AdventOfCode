package utils

import (
	"sync"
)

type Stack struct {
	lock sync.Mutex
	s    []string
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]string, 0)}
}

func (s *Stack) Push(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *Stack) Pop() (string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	// if l == 0 {
	// 	return "", errors.New("Empty Stack")
	// }

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}
