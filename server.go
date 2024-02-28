package main

import (
	"bufio"
	"fmt"
	"log"
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

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func server(port1, port2 string) {
	fmt.Println("Launching server...")
	fmt.Println("Server has been launched on " + GetLocalIP().String())

	// Устанавливаем прослушивание порта
	port, err := net.Listen("tcp", "127.0.0.1:"+port1)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Открываем порт
	client, err := port.Accept()
	if err != nil {
		log.Println(err.Error())
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go connConversation(conn1, conn2)
	go connConversation(conn2, conn1)
	wg.Wait()
}
