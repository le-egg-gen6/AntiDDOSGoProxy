package command

import "net"

func Hello(conn net.Conn, args []string) string {
	return "Hello"
}
