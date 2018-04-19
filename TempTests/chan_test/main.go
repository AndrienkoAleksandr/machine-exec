package main

import (
	"time"
	"fmt"
)

func main()  {
	stringChan := make(chan string)

	go func() {
		stringChan <- "test1"

		timer := time.NewTimer(1 * time.Second)
		<-timer.C

		stringChan <- "test2"

		timer2 := time.NewTimer(1 * time.Second)
		<-timer2.C

		stringChan <- "test3"

	}()

	go func() {
		for  {
			line := <- stringChan
			fmt.Println("first awaiter " + line)
		}
	}()

	go func() {
		for  {
			line := <- stringChan
			fmt.Println("first awaiter 2 " + line)
		}
	}()

	timer := time.NewTimer(6 * time.Second)
	<-timer.C
	fmt.Println("Timer expired")
}
