package telnet

import "net"

type CommandHandler func(conn net.Conn, args []string) string

type CommandManager struct {
	handlers map[string]CommandHandler
}

func NewCommandManager() *CommandManager {
	return &CommandManager{
		handlers: make(map[string]CommandHandler),
	}
}

func (cm *CommandManager) Register(name string, handler CommandHandler) {
	cm.handlers[name] = handler
}
