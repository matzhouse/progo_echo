package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
    "strconv"
)

const (
	PORT = 8080
)

func handleConn(client net.Conn) {
    b := bufio.NewReader(client)
    for {
        line, err := b.ReadBytes('\n')
        if err != nil { // EOF, or worse
            break
        }
        client.Write(line)
    }
}

func main() {

	fmt.Println("listening..")
	listener, err := net.Listen("tcp", "0.0.0.0:" + strconv.Itoa(PORT))
	if err != nil {
		println("error listening:", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			println("Error accept:", err.Error())
			return
		}
		handleConn(conn)
	}

}