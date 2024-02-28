package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func listener(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
	}
}

func writer(conn net.Conn, mes chan string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "exit" {
			mes <- text
			return
		}
		fmt.Fprintf(conn, text+"\n")
	}
}

func client(host, port, text string) {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(conn, text+"\n")
	var exit chan string
	go listener(conn)
	go writer(conn, exit)
	<-exit
	//for {
	//	conn.Write([]byte(text + "\n"))
	//	time.Sleep(time.Second)
	//}
}
