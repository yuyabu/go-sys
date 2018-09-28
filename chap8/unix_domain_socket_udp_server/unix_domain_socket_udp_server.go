package unix_domain_socket_udp

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-server")
	//エラーチェックは削除(存在しなかったらしなかったで問題ないので不要)
	os.Remove(path)
	fmt.Println("server is running at " + path)
	conn, err := net.ListenPacket("unixgram", path)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v\n",
			remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)

		if err != nil {
			panic(err)
		}
	}
}
