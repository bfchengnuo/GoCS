package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 模拟 nc
// 从建立的连接中读取数据，写出到标准输出
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
