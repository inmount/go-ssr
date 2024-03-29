package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

var working bool

// 错误处理
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error :", err.Error())
		os.Exit(1)
	}
}

// 接受连接线程
func acceptor(conn *net.TCPConn) {

	// 约定退出事件
	defer conn.Close()

	// 获取客户端IP地址。
	var remoteAddr = conn.RemoteAddr()
	fmt.Println("-> Accept：", remoteAddr)
	fmt.Println("[", remoteAddr, "] > Reading ...")

	// 定义缓存
	buffer := make([]byte, 4096)

	for working {

		// 读取内容。
		// var bs, err = ioutil.ReadAll(conn)
		re, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("[", remoteAddr, "] > Error：", err)
			break
		}

		// 判断数组长度
		if re > 0 {
			// 显示字符串
			fmt.Println("[", remoteAddr, "] > Read string：", string(buffer))

			// 原封不动的返回
			conn.Write(buffer)
		}

		// 临时让出控制权
		runtime.Gosched()

	}

}

func main() {

	// 定义端口
	local, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8602")
	checkError(err)

	// 开始监听
	listener, err := net.ListenTCP("tcp", local)
	checkError(err)
	working = true

	// 约定关闭监听
	defer func() {
		listener.Close()
		working = false
	}()
	fmt.Println("-> Server is start.")

	// 循环接受连接
	for working {
		// 接受连接
		var accept, errAccept = listener.AcceptTCP()
		if errAccept != nil {
			fmt.Println("-> Accept error：", errAccept)
			return
		}
		// 建立数据处理线程
		go acceptor(accept)
	}
}
