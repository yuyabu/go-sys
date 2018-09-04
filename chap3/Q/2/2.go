//Q3.2テスト用の適当なサイズのファイルを作成
package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	io.CopyN(file, rand.Reader, 1024)

}
