package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("プロセスID:%d\n", os.Getpid())
	fmt.Printf("親プロセスID:%d\n", os.Getppid())
	fmt.Printf("user_id:%d\n", os.Geteuid())
	fmt.Printf("group_id:%d\n", os.Getegid())
	fmt.Printf("\n")
	for {
		fmt.Println(".")
		time.Sleep(1 * time.Second)
	}
}
