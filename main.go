package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	ip := flag.String("ip", "localhost", "Server IP to connect to")
	port := flag.Int("port", 1339, "Port to use on the server")
	address := *ip + ":" + strconv.Itoa(*port)

	fmt.Println("Connecting to", address + "...")
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error: could not connect to the server")
		return
	}

	fmt.Println("Connected to", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	fmt.Println("Staring game...")
	_, err = conn.Write([]byte("START\n"))
	
	if err != nil {
		fmt.Println("Error: could not connect to the server")
		return
	}

	fmt.Println("Exiting...")
	conn.Write([]byte("EXIT\n")) // Why would we check the answer?
	return
}

func ParseAnswer(r *bufio.Reader) (ok bool, answer string, err error) {
	response, err := r.ReadString('\n')
	resp := strings.Split(response, " ")
	if resp[0] == "OK" {
		ok = true
	} else if resp[0] == "ERR" {
		ok = false
	}
	answer = strings.Join(resp[1:], " ")
	return
}

