package telnet

import (
	"AntiDDOSGoProxy/telnet/command"
	"bufio"
	"fmt"
	"net"
	"strings"
)

const (
	Host             = "0.0.0.0"
	Port             = "23333"
	MaxLoginAttempts = 3
)

var AllowedIPs = []string{"127.0.0.1", "::1", "27.128.12.34", "138.128.112.12"}

var users = map[string]string{
	"admin":     "admin",
	"ledeptrai": "ledeptrai",
}

var CmdManagerInstance *CommandManager

func sendMessage(conn net.Conn, message string, newLine bool) {
	if newLine {
		conn.Write([]byte(message + "\r\n"))
	} else {
		conn.Write([]byte(message))
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("New client connected: %s\n", clientAddr)

	clientIP := conn.RemoteAddr().(*net.TCPAddr).IP.String()
	if !isAllowedIP(clientIP) {
		fmt.Printf("Blocked unauthorized IP: %s\n", clientIP)
		sendMessage(conn, "Access denied.", true)
		return
	}

	// Authentication process
	sendMessage(conn, "Welcome to Go Telnet Server!", true)
	sendMessage(conn, "Please log in.", true)

	authenticated := false
	var username string
	scanner := bufio.NewScanner(conn)

	for attempts := 0; attempts < MaxLoginAttempts; attempts++ {
		sendMessage(conn, "Username: ", false)
		if !scanner.Scan() {
			return
		}
		username = strings.TrimSpace(scanner.Text())

		sendMessage(conn, "Password: ", false)
		if !scanner.Scan() {
			return
		}
		password := strings.TrimSpace(scanner.Text())

		// Check credentials
		if validPassword, exists := users[username]; exists && password == validPassword {
			authenticated = true
			break
		} else {
			sendMessage(conn, "Invalid credentials. Try again.", true)
		}
	}

	if !authenticated {
		sendMessage(conn, "Too many failed attempts. Disconnecting.", true)
		fmt.Printf("Client %s failed login attempts and was disconnected.\n", clientAddr)
		return
	}

	sendMessage(conn, "Login successful!", true)
	sendMessage(conn, "Exit command: exit", true)
	fmt.Printf("User %s authenticated from %s\n", username, clientAddr)

	// Start session after successful login
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		command_ := args[0]
		args = args[1:] // Remove the command_ from args list

		// Execute registered command_
		if handler, found := CmdManagerInstance.handlers[command_]; found {
			sendMessage(conn, fmt.Sprintf("Command result: %s", handler(conn, args)), true)
		} else {
			// If command_ was "exit", break out of loop
			if command_ == "exit" {
				fmt.Printf("User %s disconnected.\n", username)
				break
			}
			sendMessage(conn, "Unknown command.", true)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from client %s: %v\n", clientAddr, err)
	}

}

func isAllowedIP(ip string) bool {
	for _, allowedIP := range AllowedIPs {
		if ip == allowedIP {
			return true
		}
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP.IsLoopback() || parsedIP.IsPrivate() {
		return true
	}

	return false
}

func StartTelnetServer() {
	listener, err := net.Listen("tcp", Host+":"+Port)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	CmdManagerInstance = NewCommandManager()

	// register handler
	CmdManagerInstance.Register("hello", command.Hello)

	for {
		conn, err := listener.Accept()
		if err != nil {
			//
			continue
		}
		go handleConnection(conn)
	}
}
