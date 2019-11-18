// goroutines represents concurrent tasks, basically a thread
// channels are used to communicate between tasks
// select enables task synchronization
// Goroutines are functions or methods that run concurrently with other functions or methods.
package concurrency

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello")
}

func Perform_goroutine() {
	go hello()
	fmt.Println("perform_goroutine function")

	Using_channels()
}