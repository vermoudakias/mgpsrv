package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

type Worker struct {
	req chan int
	res chan int
}

var MAX_WORKERS = 4
var WORKER_STOP = 0
var WORKER_STOPPED = 0

func main() {
	fmt.Println("Server: Starting")
	// Setup channel on which to send signal notifications
	chan_kill := make(chan os.Signal, 1)
	signal.Notify(chan_kill, os.Interrupt, os.Kill)
	// Start workers
	workers := make([]WorkerContext, MAX_WORKERS)
	for i := 0; i < MAX_WORKERS; i++ {
		workers[i].req = make(chan int, 1)
		workers[i].res = make(chan int, 1)
		go WorkerProcess(&workers[i])
	}
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Server: Timeout")
	case <-chan_kill:
		fmt.Println("Server: Interrupted")
	}
	// Stop workers
	for i := 0; i < MAX_WORKERS; i++ {
		fmt.Println("Server: Notify worker to stop", i)
		workers[i].req <- WORKER_STOP
		<-workers[i].res
	}
	fmt.Println("Server: Stopped")
}
