package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

func connConversation(conn1, conn2 net.Conn) {
	for {
		// Будем прослушивать все сообщения разделенные \n
		message, _ := bufio.NewReader(conn1).ReadString('\n')
		// Распечатываем полученое сообщение
		fmt.Print("Message Received:", string(message))
		// Процесс выборки для полученной строки
		newmessage := strings.ToUpper(message)
		// Отправить новую строку обратно клиенту
		conn2.Write([]byte(newmessage + "\n"))
	}
}

func server(port1, port2 string) {
	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	client1, _ := net.Listen("tcp", ":"+port1)
	client2, _ := net.Listen("tcp", ":"+port2)

	// Открываем порт
	conn1, _ := client1.Accept()
	conn2, _ := client2.Accept()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go connConversation(conn1, conn2)
	go connConversation(conn2, conn1)
	wg.Wait()
}
