package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// 获取本机MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}

	for _, inter := range interfaces {
		mac := inter.HardwareAddr
		fmt.Println("MAC = ", mac)
	}

	// 获取本机IP
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(addrs)

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		fmt.Println("*******************************")
		fmt.Println(address)
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

}
