//Channels
/********************************************************************************************************
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

ch := make(chan int)
By default, sends and receives block until the other side is ready.
This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines.
Once both goroutines have completed their computation, it calculates the final result.
********************************************************************************************************/
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	fmt.Println("Length of s array: ", len(s))
	// create channel c
	c := make(chan int)
	// goroutine 1
	go sum(s[:len(s)/2], c)
	// slice [0:3]
	fmt.Println("Slice 1: ", s[:len(s)/2])
	// goroutine 2
	go sum(s[len(s)/2:], c)
	// slice [3:]
	fmt.Println("Slice 2: ", s[len(s)/2:])
	// channel receiver
	x, y := <-c, <-c // receive from c

	fmt.Println("Channel receiver 1 value: ", x)
	fmt.Println("Channel receiver 2 value: ", y)
	fmt.Println("Add: ", x+y)
}
