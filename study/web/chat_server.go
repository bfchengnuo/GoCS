package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 简单的聊天程序服务端，参考并发内容
// 客户端使用 nc 或者 tcp_client.go
func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	// 接收客户端连接与断开、消息分发事件
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

// 单向输入
type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// 分发给所有客户端
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			// 用户登陆
			clients[cli] = true
		case cli := <-leaving:
			// 用户离开
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "you are " + who
	messages <- who + " 上线."
	entering <- ch

	// 获取输入
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + " : " + input.Text()
	}

	// 离开
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

// 打印接收到的消息
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range <-ch {
		fmt.Fprintln(conn, msg)
	}
}
