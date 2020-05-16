package main

import (
	"io"
	"log"
	"net"
	"time"
)

// tcp 服务模拟
// linux 下可以使用 nc 测试
func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 阻塞
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		// 不停的往连接中输出
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
