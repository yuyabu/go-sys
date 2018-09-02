package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	//json化する元のデータ
	source := map[string]string{
		"hello": "world",
	}
	//ここにコードを書く

	writer := gzip.NewWriter(w)
	writer.Header.Name = "testttt"
	mlutiWriter := io.MultiWriter(writer, os.Stdout)
	encoder := json.NewEncoder(mlutiWriter)
	encoder.SetIndent("", "   ")
	encoder.Encode(source)
	writer.Flush()
	writer.Close()
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
