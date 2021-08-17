package main

import (
	"fmt"

	"github.com/patelpreet422/wifi-file-transfer/util"
)

func main() {
	ip, _ := util.GetIPAddr()
	fmt.Println("ip: %v\n", ip)
}
