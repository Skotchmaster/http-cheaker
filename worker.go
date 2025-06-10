package main

import (
	"fmt"
	"net/http"
	"sync"
)

func NewPool(bufferSize int) *Pool {
	return &Pool{
		jobs:    make(chan string, bufferSize),
		quit:    make(chan struct{}),
		workers: make(map[int]struct{}),
	}
}

func (p *Pool) AddWorker() {
	p.mu.Lock()
	id := p.nextID
	p.nextID++
	p.workers[id] = struct{}{}
	p.mu.Unlock()

	p.wg.Add(1)
	go worker(id, p.jobs, p.quit, &p.wg)
	fmt.Printf("Added worker #%d (total: %d)\n", id, p.Status())
}

func (p *Pool) RemoveWorker() {
	p.mu.Lock()
	var id int
	for i := range p.workers {
		id = i
		break
	}
	delete(p.workers, id)
	p.mu.Unlock()

	p.quit <- struct{}{}
	fmt.Printf("Removed worker #%d (remaining: %d)\n", id, p.Status())
}

func (p *Pool) Submit(job string) {
	p.jobs <- job
}

func (p *Pool) Status() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.workers)
}

func (p *Pool) Shutdown() {
	count := p.Status()
	for i := 0; i < count; i++ {
		p.quit <- struct{}{}
	}
	p.wg.Wait()
	close(p.jobs)
	close(p.quit)
}

func worker(id int, jobs <-chan string, quit <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case url := <-jobs:
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("[Worker #%d] ERROR %s → %v\n", id, url, err)
			} else {
				fmt.Printf("[Worker #%d] %s → %s\n", id, url, resp.Status)
				resp.Body.Close()
			}
		case <-quit:
			fmt.Printf("[Worker #%d] shutting down\n", id)
			return
		}
	}
}
