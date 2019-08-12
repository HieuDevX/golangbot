package main

import "fmt"

// type cast to unidirectional channels, in this function, channel is unidirectional channels
func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)
}
