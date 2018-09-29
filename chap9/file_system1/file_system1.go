package main

import (
	"fmt"
	"io"
	"os"
)

//新規作成
func open() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

//読み込み
func read() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

//追記モードで開く
func append() {
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	io.WriteString(file, "Appended content\n")
}

//ディレクトリの作成
func makeDirectory() {
	os.MkdirAll("test_dir/test", 0755)
}
func main() {
	open()
	read()
	append()
	makeDirectory()
}
