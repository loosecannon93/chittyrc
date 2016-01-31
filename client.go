package main

import (
	"github.com/loosecannon93/chittyrc/lib/chilog"
	"net"
	"strings"
)

func handleClient(server *Server, conn net.Conn) {
	server.mutex.Lock()
	user := new(User)
	user.conn = conn
	server.users = append(server.users, user)
	server.mutex.Unlock()

	chilog.Info.Println("Trying to get remote hostname")
	addr, _, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		chilog.Error.Fatalln("Could not split hostname and port", conn.RemoteAddr().String())
	}

	hostnames, err := net.LookupAddr(addr)
	if err != nil {
		chilog.Error.Println("Could not get hostname: ", err)
		user.host = conn.RemoteAddr().String()
	} else {
		user.host = hostnames[0]
	}
	buffer := make([]byte, MSG_LEN)
	message := ""
	for {
		bytes_read, err := conn.Read(buffer)
		if err != nil {
			chilog.Error.Println("Error reading from client socket:", err)
			return
		}
		message += string(buffer[:bytes_read])
		chilog.Info.Printf("Got message: %q \n", message)
		length := len(message)
		if length < MSG_LEN {
			chilog.Debug.Printf("last chars: %q\n", message[length-2:length])
			if strings.HasSuffix(message, "\r\n") {
				handleMessage(server, user, message)
				message = ""
			}
		} else {
			message = message[:MSG_LEN-2]
			message += "\r\n"
			handleMessage(server, user, message)
		}

	}

}

func handleMessage(server *Server, user *User, message string) {
	tokens := strings.Split(strings.TrimSuffix(message, "\r\n"), " ")
	chilog.Debug.Printf("tokenized command %q\n", tokens)

	command := tokens[0]
	args := make([]string, 0)
	for i, arg := range tokens[1:] {
		chilog.Debug.Printf("got token %d  = %s\n", i, arg)
		if arg[0] == ':' {
			args = append(args, strings.Join(tokens[i+1:], " "))
			break
		}
		args = append(args, arg)
	}
	args_count := len(args)
	args[args_count-1] = strings.TrimPrefix(args[args_count-1], ":")

	chilog.Debug.Printf("tokenized arguments %q\n", args)
	handlers[command](server, user, args)
}
