package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func startServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server started. Listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	syn, _ := reader.ReadString('\n')
	if syn == "SYN\n" {
		fmt.Println("Server received SYN from the client")
		time.Sleep(1 *  time.Second)

		fmt.Println("Server sending SYN-ACK to the client")
		conn.Write([]byte("SYN-ACK\n"))

		ack, _ := reader.ReadString('\n')
		if ack == "ACK\n" {
			fmt.Println("Server received ACK from the client")
		} else {
			fmt.Println("Invalid response from the clinet: ", ack)
		}
	}
}

func startClient() {
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error dialing: ", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Client sending SYN to the server")
	fmt.Fprintf(conn, "SYN\n")

	reader := bufio.NewReader(conn)
	synAck, _ := reader.ReadString('\n')
	if synAck == "SYN-ACK\n" {
		fmt.Println("Client received SYN-ACK from the server")
		time.Sleep(1 * time.Second)

		fmt.Println("Client sending ACK to the server")
		fmt.Fprintf(conn, "ACK\n")

		fmt.Println("Handshake complete. Connection established")
	} else {
		fmt.Println("Invalid response from the server: ", synAck)
	}
}

func main() {
	go startServer()
	go startClient()

	time.Sleep(5 * time.Second)
}