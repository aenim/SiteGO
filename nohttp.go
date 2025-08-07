package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print("Ошибка ", err)
			continue
		}
		go handleHTTP(conn)
	}
}

func handleHTTP(conn net.Conn) {
	defer conn.Close()                // отложенная инструкция на закрытие соединения
	reader := bufio.NewReader(conn)   // чтение запросов в буфер
	tp := textproto.NewReader(reader) //

	// Читаем первую строку (Request-Line)
	requestLine, err := tp.ReadLine()
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	fmt.Print(requestLine)

	parts := strings.Split(requestLine, " ") // разделяем запрос по пробелам
	if len(parts) != 3 {
		fmt.Println("Malformed request line") // если в запросе не 3 слова после сплита значит чтото не так
		return
	}

	method, path, version := parts[0], parts[1], parts[2] //
	fmt.Printf("Method: %s, Path: %s, Version: %s\n", method, path, version)

	header, err := tp.ReadMIMEHeader() // чтение заголовка
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Headers ", header)

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Connection: close\r\n" +
		"\r\n" +
		fmt.Sprintf("You requested %s %s\n", method, path)

	conn.Write([]byte(response))

}

// добавить обработку соединения гарантированное закрытие сессии ОК
// буферизированный читатель ОК
// читатель НТТР заголовков
// чтение зароса
// Читаем первую строку HTTP-запроса, которая имеет формат: МЕТОД URI ВЕРСИЯ_HTTP
// Разбиваем строку запроса на части.
// чтение заголовков
// формирование и отправка ответа
