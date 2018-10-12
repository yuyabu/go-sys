package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//サイズが1より大きいチャネルを作成
	signals := make(chan os.Signal, 1)

	//最初のチャネル以降は、可変長引数で任意の数のシグナルを設定可能
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	s := <-signals

	switch s {
	case syscall.SIGINT:
		fmt.Println("SIGINT")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM")
	}
}
