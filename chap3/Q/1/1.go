/*
古いファイルを新しいフィあるにコピーする*/
package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("old.txt")

	if err != nil {
		panic(err)
	}
	file2, err := os.Create("new.txt")

	if err != nil {
		panic(err)
	}

	//fmt.Println(file)
	io.Copy(file2, file)

	//writeSize,err:=ioCopy(writer,reader)
}
