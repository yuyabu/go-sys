//io.Writeのデコレータを使ってgzipの圧縮を使うやつを手本に(2.45)
//encoding/csvというパッケージを使って標準入出力やファイルにcsvを出力してみよう
package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// file, err := os.Create("test.txt.gz")
	// if err != nil {
	// 	panic(err)
	// }

	// writer := gzip.NewWriter(file)
	// writer.Header.Name = "test.txt"
	// io.WriteString(writer, "gzip.Writer example\n")
	// writer.Close()
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"val1", "val2"})
	writer.Flush()
	writer.Write([]string{"val3", "val4"})
	writer.Flush()

}
