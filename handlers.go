package main

import (
	"fmt"
	"github.com/loosecannon93/chittyrc/lib/chilog"
	//	"net"
	//	"strings"
)

type handler func(*Server, *User, []string)

var handlers = map[string]handler{
	"NICK": handle_NICK,
	"USER": handle_USER,
}

func handle_NICK(server *Server, user *User, args []string) {
	chilog.Info.Println("Registering User with NICK:", args[0])
	user.nick = args[0]
	if user.name != "" && user.nick != "" {
		send_RPL_WELCOME(server, user)
	}
}

func handle_USER(server *Server, user *User, args []string) {
	chilog.Info.Println("Registering User with USER:", args[0], "fullname:", args[3])
	name := args[0]
	full_name := args[3]
	user.name = name
	user.full_name = full_name

	if user.name != "" && user.nick != "" {
		send_RPL_WELCOME(server, user)
	}
}

func send_RPL_WELCOME(server *Server, user *User) {
	reply := fmt.Sprintf(":%s %s %s :Welcome to the Internet Relay Network %s!%s@%s\r\n",
		server.host, RPL_WELCOME, user.nick, user.nick, user.name, user.host)
	send_reply(server, user, reply)
}

func send_reply(server *Server, user *User, message string) {
	buffer := []byte(message)
	bytes_to_send := len(buffer)
	bytes_sent := 0
	for bytes_sent < bytes_to_send {
		bytes_sent, _ = user.conn.Write(buffer[bytes_sent:])
	}
}
