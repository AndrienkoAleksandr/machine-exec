package main

import (
	"time"
	"fmt"
)

func main()  {

	withGoRotines()

	timer := time.NewTimer(6 * time.Second)
	<-timer.C
	fmt.Println("Timer expired")
}

func withGoRotines() {
	defer fmt.Println("Closed")
	go func() {
		for {
			timer := time.NewTimer(1 * time.Second)
			<-timer.C
			fmt.Println("time out")
		}
	}()
}
