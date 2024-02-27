package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
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

func client(port string) {
	conn, _ := net.Dial("tcp", "127.0.0.1:"+port)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go listener(conn)
	go writer(conn)
	wg.Wait()
}
