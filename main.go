package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	initialWorkers := flag.Int("workers", 5, "initial number of workers in pool")
	flag.Parse()

	pool := NewPool(100)
	for i := 0; i < *initialWorkers; i++ {
		pool.AddWorker()
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите URL или команду (add, remove, status, quit):")

	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "add":
			pool.AddWorker()
		case "remove":
			pool.RemoveWorker()
		case "status":
			fmt.Printf("Active workers: %d\n", pool.Status())
		case "quit":
			pool.Shutdown()
			os.Exit(0)
		default:
			pool.Submit(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
