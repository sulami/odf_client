package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	Conn net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func (s *Server) Connect(addr string) (err error) {
	fmt.Println("Connecting to", addr)
	s.Conn, err = net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error: could not connect to the server")
		return
	} else {
		fmt.Println("Connected to", s.Conn.RemoteAddr())
	}
	s.Reader = bufio.NewReader(s.Conn)
	s.Writer = bufio.NewWriter(s.Conn)

	go s.Read()

	return
}

func (s *Server) Write(msg string) {
	s.Writer.WriteString(msg + "\n")
	s.Writer.Flush()
}

func (s *Server) Read() {
	for {
		response, err := s.Reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		resp := strings.Split(response, " ")
		if resp[0] == "OK" {
			fmt.Println(strings.Join(resp[1:], " "))
		} else if resp[0] == "ERR" {
			fmt.Println("Error:", resp[1])
		}
	}
}

func Read(s *Server) {
	for {
		var buf string
		fmt.Scanln(&buf)
		s.Write(strings.ToUpper(buf))
		if strings.ToUpper(buf) == "EXIT" {
			break
		}
	}
}

func main() {
	ip := flag.String("ip", "localhost", "Server IP to connect to")
	port := flag.Int("port", 1339, "Port to use on the server")
	flag.Parse()

	server := &Server{}
	err := server.Connect(*ip + ":" + strconv.Itoa(*port))
	if err != nil {
		return
	}

	Read(server)
}

