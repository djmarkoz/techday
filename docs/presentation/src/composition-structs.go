package main

import "sync"

type Sum struct {
	sync.RWMutex
	i int
}

func (s *Sum) Get() int {
	s.Lock()
	defer s.Unlock()

	return s.i
}

func (s *Sum) Increase() int {
	s.Lock()
	defer s.Unlock()

	s.i++
	return s.i
}
