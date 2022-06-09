package main

import (
	"fmt"
	"sync"
	"time"
)

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func (s *Service) Work(job int) {
	s.Lock.RLock()
	exist := s.InProgress[job]
	if exist {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Waiting for response job %d\n", job)
		res := <-response
		fmt.Printf("Response done, received %d\n", res)
		return
	}
	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()
	fmt.Printf("calculating fibo for %d\n", job)
	result := ExpensiveFibonacci(job)

	s.Lock.RLock()
	pendWorkers, exists := s.IsPending[job]
	s.Lock.RUnlock()
	if exists {
		for _, pendWorker := range pendWorkers {
			pendWorker <- result
		}
		fmt.Printf("Result sent - all pending workers ready, job: %d\n", job)
	}
	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func ExpensiveFibonacci(n int) int {
	fmt.Printf("Calculating expensive fibo for %d\n", n)
	time.Sleep(5 * time.Second)
	return n
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 5, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	wg.Wait()
}
