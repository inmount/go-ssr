package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {

	// 定义端口
	local, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8602")

	// 开始监听
	var listener, errListen = net.ListenTCP("tcp", local)
	if errListen != nil {
		fmt.Println("-> Listen error：", errListen)
		return
	}

	// 约定关闭监听
	defer func() {
		listener.Close()
	}()
	fmt.Println("-> Server is start.")

	// 接受连接
	var accept, errAccept = listener.AcceptTCP()
	if errAccept != nil {
		fmt.Println("-> Accept error：", errAccept)
		return
	}

	// 获取客户端IP地址。
	var remoteAddr = accept.RemoteAddr()
	fmt.Println("-> Accept：", remoteAddr)
	fmt.Println("-> Reading ...")

	// 读取内容。
	var bs, _ = ioutil.ReadAll(accept)
	fmt.Println("-> Read string：", string(bs))

	// 原封不动的返回
	accept.Write(bs)

	// 关闭连接
	accept.Close()
}
