// A channel is a pipe between goroutines to synchronize execution and communicate by sending/receiving data.

package main

import "fmt"

func sendValues(myIntChannel chan int) {
	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}
	close(myIntChannel) // close() is always called from the sender
}

func main() {
	// generalChannel()
	bufferedChannel()
}

func generalChannel() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 7; i++ {
		fmt.Println(<-myIntChannel) //receiving value
		// Reading from a closed channel is also allowed !! last two values will be 0 here.

		// We cloud read like below to detect if the channel is closed :
		// value, open := <-myIntChannel
		// if !open {
		//    break;
		// }

		// exact same functionality will be achieved by ranging on a loop
		// for value := range myIntChannel { ... }
	}
}

func bufferedChannel() {
	mychannel := make(chan int, 2)
	mychannel <- 10 // this will be a blocking call for general channel, but not for a buffered channel !!
	fmt.Println(<-mychannel)

	// we can say that if there are no receive operations for a channel, a goroutine can still perform c sending operations,
	// where c is the capacity of the buffered channel.
}
