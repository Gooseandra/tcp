package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func connConversation(conn1 net.Conn) {
	for {
		// Будем прослушивать все сообщения разделенные \n
		message, err := bufio.NewReader(conn1).ReadString('\n')
		if err != nil {
			log.Println(err)
			conn1.Close()
			return
		} // Распечатываем полученое сообщение
		fmt.Println("Message Received:", message)
		// Процесс выборки для полученной строки
		newmessage := strings.ToUpper(message)
		// Отправить новую строку обратно клиенту
		conn1.Write([]byte(newmessage + "\n"))
	}
}

func server(port1 string) {
	fmt.Println("Launching server...")

	listen, err := net.Listen("tcp", ":"+port1)
	if err != nil {
		log.Println(err)
		return
	} // Открываем порт
	for {
		client, err := listen.Accept()
		if err != nil {
			log.Println(err)
			return
		} // Открываем порт
		log.Println("new client")
		go connConversation(client)
	}

}
