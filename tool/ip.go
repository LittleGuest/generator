package tool

import (
	"log"
	"net"
)

// 获取IP地址（IPv4）
func GetIPv4() net.IP {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, v := range interfaces {
		if v.Flags&net.FlagUp == 0 {
			continue
		}
		if v.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := v.Addrs()
		if err != nil {
			log.Println(err)
			return nil
		}
		for _, addr := range addrs {
			ip := net.IP{}
			switch vv := addr.(type) {
			case *net.IPNet:
				ip = vv.IP
			case *net.IPAddr:
				ip = vv.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip
		}
	}
	return nil
}
