//zipファイルの書き込み
package main

import (
	"archive/zip"
	"os"
)

func main() {
	file, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// writer, err := zipWriter.Create("newfile.txt")
	// //test := strings.NewReader("TEST")
	// writer.(test)
}
