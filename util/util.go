package util

import (
	"errors"
	"net"
)

func GetIPAddr() (string, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return "", errors.New("util.getIPAddr(): Failed to get net interfaces")
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return "", errors.New("util.getIPAddr(): Failed to get IP addresses of interface")
		}

		// Each interface can have both IPv4 and IPv6
		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			}

			if !ip.IsLoopback() {
				return ip.String(), nil
			}
		}
	}
	return "", errors.New("util.getIPAddr(): No local IP found")

}
