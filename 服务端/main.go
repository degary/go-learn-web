package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server")
	listen, err := net.Listen("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {

	defer conn.Close()
	for {
		//定义数组,接收数据
		var buf [128]byte
		//把数组转换成切片
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read failed ,err:", err)
			break
		}
		str := string(buf[:n])
		fmt.Println("recv from client,data:", str)
	}
}
