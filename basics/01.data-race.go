//A data race happens when processes have to access the same variable concurrently i.e. one process reads from a memory location while another simultaneously writes to the exact same memory location.
//
//A race condition is a flaw in a program regarding the timing/ordering of operations which disrupts the logic of the program and leads to erroneous results.

package main

import "fmt"

func deposit(balance *int, amount int) {
	*balance += amount
}

func withdraw(balance *int, amount int) {
	*balance -= amount
}

func main() {

	balance := 100

	go deposit(&balance, 10) //depositing 10

	withdraw(&balance, 50) //withdrawing 50

	fmt.Println(balance)
}

// [golang official race-detector](https://go.dev/doc/articles/race_detector) holds some good examples of data races.
