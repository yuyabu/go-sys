package main

import (
	"fmt"
	"time"
)

func main() {
	l := NewFileLock("main.go")
	fmt.Println("try locking ...")
	l.Lock()
	fmt.Println("locked")
	time.Sleep(10 * time.Second)
	l.Unlock()
	fmt.Println("unlock")
}
