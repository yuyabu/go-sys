package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader
	done := make(chan bool)
	//ここにIOパッケージを使ったコードをかく
	var a io.Reader
	// var s io.Reader
	// var c io.Reader
	// var i io.Reader

	//シークしてio.CopyNをしまくればいけると思う。
	// var writer io.Writer
	go func() {
		programming.Seek(5, 0)
		a = io.LimitReader(programming, 1)
		done <- true
	}()
	<-done
	stream = io.MultiReader(a)
	io.Copy(os.Stdout, stream)

}
