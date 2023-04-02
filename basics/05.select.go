// The select statement blocks the code and waits for multiple channel operations simultaneously.
// select {
//  case channel operation:
//    ...
//  more cases :
//    ...
//  default : //Optional
//    ...
//}

package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel1 <- "I'll print every 100ms"
			time.Sleep(time.Millisecond * 100)

		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel2 <- "I'll print every 300ms"
			time.Sleep(time.Millisecond * 300)

		}
	}()

	for i := 0; i < 5; i++ {
		//fmt.Println(<-channel1)
		//fmt.Println(<-channel2)

		// above two prints will print 10 lines (5 of each) by turn
		// to just select random 5 lines ::

		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		}

	}
}
