package SRP_net

import "net"

func ListenerHandler(netListener net.Listener) {
	defer netListener.Close()
	for {
		conn, _ := netListener.Accept()
		go connectionHandler(conn)
	}
}
