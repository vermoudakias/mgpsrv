package main

import (
	"fmt"
)

type WorkerContext struct {
	req chan int
	res chan int
}

func WorkerProcess(ctx *WorkerContext) {
	var rq, rs int
	finished := false
	fmt.Println("Worker: Starting")
	for finished != true {
		select {
		case rq = <-ctx.req:
			fmt.Println("Worker: Request received")
			rs, finished = handleReq(rq)
			ctx.res <- rs
		}
	}
	fmt.Println("Worker: Stopped")
}

func handleReq(r int) (res int, finished bool) {
	if r == 0 {
		finished = true
	} else {
		finished = false
	}
	res = r
	return
}
