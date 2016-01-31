package main

import (
	"flag"
	"github.com/loosecannon93/chittyrc/lib/chilog"
	"net"
	"os"
)

func main() {

	// Parse Command line flags
	var (
		oper_password string
		port          string
		verbose       bool
		vverbose      bool
		quiet         bool
	)

	flag.StringVar(&oper_password, "o", "pass", "Operator Password")
	flag.BoolVar(&verbose, "v", false, "Verbose: log upto DEBUG level")
	flag.BoolVar(&vverbose, "vv", false, "Very Verbose: log everything including TRACE")
	flag.BoolVar(&quiet, "q", false, "Print nothing to the log")
	flag.StringVar(&port, "p", "6667", "Port to bind to")

	flag.Parse()

	// Set the log level
	log_level := chilog.INFO

	switch {
	case quiet:
		log_level = chilog.QUIET
	case vverbose:
		log_level = chilog.TRACE
	case verbose:
		log_level = chilog.DEBUG
	}
	chilog.Init(log_level)

	chilog.Info.Println("Server Starting")
	chilog.Info.Print("Operator password is '", oper_password, "'")
	chilog.Info.Println("port is", port)

	socket, err := net.Listen("tcp", ":"+port)
	if err != nil {
		chilog.Critical.Fatalln("Server failed to bind to port ", port)
	}
	server := InitServer()
	server.host, _ = os.Hostname()
	for {
		conn, err := socket.Accept()
		if err != nil {
			chilog.Critical.Fatalln("Failed to Accept() on socket")
		}
		go handleClient(server, conn)
	}
}
