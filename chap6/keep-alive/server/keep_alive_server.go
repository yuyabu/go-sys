package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			fmt.Println("Accept %V\n", conn.RemoteAddr())

			//Accept後のソケットでなんども応答を返すためにループ
			for {
				//タイムアウトを設定
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				//リクエストを読み込む
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					//タイムアウトもしくはソケットクローズじは終了
					//それ以外はエラーにする
					neterr, ok := err.(net.Error) //ダウンキャスト
					if ok && neterr.Timeout() {
						fmt.Println("timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}
				//リクエストを表示
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				content := "Hello World\n"

				//レスポンスを書き込む
				//HTTP/1.1かつ,ContentLengthの設定が必要
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body: ioutil.NopCloser(
						strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}
