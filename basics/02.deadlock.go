// A deadlock occurs when all processes are blocked while waiting for each other and the program cannot proceed further.

// Coffman conditions (all 4 have to be true to occur a deadlock):
// 1. Mutual exclusion (at least one non-sharable resource is held by process-1)
// 2. Hold & wait (it holds the resource & wait for other resource)
// 3. No preemption (system cant take away the held resource)
// 4. Circular wait

package main

import "fmt"

func main() {
	mychannel1 := make(chan int)
	mychannel2 := make(chan int)
	mychannel3 := make(chan int)
	go func() {
		<-mychannel1
	}()

	go func() {
		mychannel2 <- 20
	}()

	go func() {
		<-mychannel3
	}()

	fmt.Println(<-mychannel2)
	// the deadlocks for 1st & 3rd channel will not be detected by go compiler.
	// It can only detect if the program is stuck as a whole
}

// Starvation happens when a process is deprived of necessary resources and is unable to complete its function.
