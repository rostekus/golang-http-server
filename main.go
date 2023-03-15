package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	fmt.Println("Server is running on " + HOST + ":" + PORT)
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// incoming request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Request: ", string(buffer[:]))

	f, err := os.OpenFile("static/index.html", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// get file content
	fileContent := make([]byte, 1024)
	_, err = f.Read(fileContent)
	if err != nil {
		log.Fatal(err)
	}
	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Lenght:%d \r\n\r\n %s", len(string(fileContent[:])), string(fileContent[:]))
	conn.Write([]byte(response))

	conn.Close()
	fmt.Println("Connection closed")
}
