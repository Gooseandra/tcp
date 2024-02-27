package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func listener(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

func writer(conn net.Conn) {
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
	}
}

func client(host, port, text string) {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn.Write([]byte(text + "\n"))
		time.Sleep(time.Second)
	}
}
