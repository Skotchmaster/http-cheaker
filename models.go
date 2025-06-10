package main

import (
	"sync"
)

type Pool struct {
	jobs    chan string
	quit    chan struct{}
	wg      sync.WaitGroup
	mu      sync.Mutex
	nextID  int
	workers map[int]struct{}
}
