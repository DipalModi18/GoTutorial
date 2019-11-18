// Channels are the pipes that connect concurrent goroutines. 
// You can send values into channels from one goroutine and receive those values into another goroutine.
package concurrency

import (
	"fmt"
)

func Using_channels() {
	// Channels are typed by the values they convey.
	messages := make(chan string)  // Create a new channel with make(chan val-type). 

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
	

	// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. 
	// Buffered channels accept a limited number of values without a corresponding receiver for those values.
	messages2 := make(chan string, 2)  // Here we make a channel of strings buffering up to 2 values.

	messages2 <- "buffered"
	messages2 <- "channel"  // Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	
	fmt.Println(<-messages2)
    fmt.Println(<-messages2)
}