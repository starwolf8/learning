package main

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("Print message from within publish function:", text)
	}()
}

func main() {
	Publish("Calling Publish", 5*time.Second)
	fmt.Println("Main 1st message")

	time.Sleep(10 * time.Second)

	fmt.Println("Complete")
}
