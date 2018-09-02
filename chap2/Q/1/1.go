package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//2.4.6の例
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())

	//演習問題："%d"数字,"%s"文字列,"%f"が浮動小数点の表示を検証する

	//2.4.1から拝借
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	//file.Write([]byte("os.File example\n"))
	fmt.Fprintf(file, "数字:%d ,文字列:%s, 浮動小数点%f", 10, "moziiii😂", 6.6)

	file.Close()
}
