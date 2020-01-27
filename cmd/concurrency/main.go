/*
CHANNELS

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
		   // assign value to v.

ch := make(chan int)

BUFFERED CHANNELS
ch := make(chan int, 100)


*/

package main

import "fmt"

//Concurrently adding numbers in a slice

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum // send sum to c
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c // receive from c

// 	fmt.Println(x, y, x+y)
// }

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10) // send max 10 times, then close
	go fibonacci(cap(c), c)
	for i := range c { // for loop recognizes channel, and will loop until closed
		fmt.Println(i)
	}
}
