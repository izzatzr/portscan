// Port Scanner in Go by Izzat Zuliya R.
// Currently this is only simple port scanner implementation, furthermore might have some cool features

package main

import (
	"net"
	"os"
)

func main() {
	if len(os.Args) <= 3 {
		println("Format: ./portscanner <host/ip> <port> <protocol>")
		os.Exit(3)
	}
	host := os.Args[1] + ":" + os.Args[2]
	proto := os.Args[3]
	sc, er := net.Dial(proto, host)
	if er != nil {
		println("Port Closed")
		println(er)
	} else {
		println("Port Opened")
		println(sc)
	}
}
