package utils

import (
	"fmt"
	"net"
	"strconv"
)

func GetIP() (string, error) {
	// 获取本地IP地址
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	var ip net.IP
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ip = ipnet.IP
			break
		}
	}
	if ip == nil {
		return "", fmt.Errorf("no valid IP address found")
	}
	return ip.String(), nil
}

func GetPort(start int) int {
	port := start // 起始端口
	for {
		ln, err := net.Listen("tcp", ":"+strconv.Itoa(start))
		if err == nil {
			ln.Close()
			break
		}
		port++
	}
	return port
}
