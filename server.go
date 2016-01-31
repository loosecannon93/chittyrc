package main

import (
	"github.com/loosecannon93/chittyrc/lib/chilog"
	"net"
	"sync"
)

type User struct {
	conn      net.Conn
	full_name string
	name      string
	nick      string
	host      string
}
type Channel struct {
	name string
}

type Server struct {
	mutex    sync.Mutex
	users    [](*User)
	channels [](*Channel)
	host     string
}

const MSG_LEN = 512

func InitServer() *Server {
	chilog.Info.Println("Created Server object. Initialize space for 32 users and channels to start")
	result := new(Server)
	result.users = make([](*User), 32)
	result.channels = make([](*Channel), 32)
	return result
}
