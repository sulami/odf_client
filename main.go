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

	fmt.Println("Logging out...")
	_, err = conn.Write([]byte("LOGOUT"))
	if err != nil {
		fmt.Println("Error: could not connect to the server")
		return
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	resp := strings.Split(response, " ")
	if resp[0] == "ERR" {
		fmt.Print("Error: Server returned ", resp[1])
	} else if resp[0] == "OK" {
		fmt.Println("Logged out")
	}
}

