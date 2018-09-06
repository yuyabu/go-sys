package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("0")

	duration := 10 * time.Second
	// waitc := make(chan time.Duration)
	<-time.After(duration)

	fmt.Printf("10")
}
