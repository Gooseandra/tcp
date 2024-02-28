package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type users struct {
	Conn net.Conn
	Name string
}

var clients []users

func connConversation(conn1 users) {
	for {
		// Будем прослушивать все сообщения разделенные \n
		message, err := bufio.NewReader(conn1.Conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			conn1.Conn.Close()
			for index, item := range clients {
				if item.Conn.RemoteAddr() == conn1.Conn.RemoteAddr() {
					clients = append(clients[:index], clients[index+1:]...)
					break
				}
			}
			return
		} // Распечатываем полученое сообщение
		fmt.Println("Message Received:", message)
		if conn1.Name == "" {
			conn1.Name = strings.TrimSuffix(message, "\n")
			conn1.Conn.Write([]byte("You registered as " + message))
			for _, item := range clients {
				if item.Conn.RemoteAddr() == conn1.Conn.RemoteAddr() {
					item.Name = strings.TrimSuffix(message, "\n") // не работает
				}
			}
		} else {
			// Процесс выборки для полученной строки
			newmessage := strings.ToUpper(message)
			// Отправить новую строку обратно клиенту
			for _, item := range clients {
				if item.Conn.RemoteAddr() != conn1.Conn.RemoteAddr() {
					item.Conn.Write([]byte(conn1.Name + ": " + newmessage + "\n"))
				}
			}
		}
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
		cl, err := listen.Accept()
		user := users{cl, ""}
		clients = append(clients, user)
		if err != nil {
			log.Println(err)
			return
		} // Открываем порт
		log.Println("new client")
		go connConversation(user)
	}
}
