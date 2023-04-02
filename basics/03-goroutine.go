// A goroutine is a function or a method which is executed concurrently with the rest of the program.

package main

import (
	"fmt"
	"time"
)

func WelcomeMessage() {
	fmt.Println("Welcome to Educative!")
}

func main() {
	go WelcomeMessage()
	go func() {
		fmt.Println("Hello World!")
	}()

	time.Sleep(time.Millisecond * 200)
}

// go uses the Fork-Join model for child go routines.
// all goroutines(of a program) have same address space
// they can be thought as lightweight threads (with some exceptions) :
//    less costly, only takes few kb
//    have their own callstacks that grow & shrink dynamically
//    no block forever
