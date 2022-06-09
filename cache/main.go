package main

import (
	"fmt"
	"log"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	result, exist := m.cache[key]
	if !exist {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}
	return result.value, result.err
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	values := []int{42, 8, 43, 42, 10}
	for _, n := range values {
		start := time.Now()
		value, err := cache.Get(n)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%d,%s,%d\n", n, time.Since(start), value)
	}
}
