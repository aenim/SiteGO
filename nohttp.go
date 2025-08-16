package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/textproto"
	"strconv"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
		//go sendCurl()

	}
}

// func sendCurl() {
// 	net.Dial("http", "localhost:8080")
// }

func handleConnection(conn net.Conn) {
	//fmt.Print(conn)
	defer conn.Close()

	reader := bufio.NewReader(conn)                            // записьв буфер
	requestline, err := textproto.NewReader(reader).ReadLine() // чтение первой строки
	if err != nil {
		fmt.Print("!!! ошибка!!! ")
	}

	parts := strings.Split(requestline, " ") // сплит первой строки по пробелу
	if len(parts) != 3 {
		fmt.Print("НЕПРАВИЛЬНЫЙ ЗАПРОС")
	}

	Method, Path, Protocol := parts[0], parts[1], parts[2]
	if 2 > 3 {
		fmt.Println(Method)
		fmt.Println(Path)
		fmt.Println(Protocol)

	}

	header, err := textproto.NewReader(reader).ReadMIMEHeader()
	if err != nil {
		fmt.Print("ошибка хедера")
	}

	// fmt.Println(header, "/n")

	// fmt.Println(header.Get("Content-Length")) // размер боди. Если не равен 0 - значит боди есть
	// fmt.Println(header.Get("Transfer-Encoding"))

	cL := header.Get("Content-Length")

	bodyLength, err := strconv.Atoi(cL)

	// fmt.Println("Это бодилентх ", bodyLength)

	var body []byte
	if bodyLength > 0 {
		body = make([]byte, bodyLength)

		// Читаем ровно bodyLength байт
		_, err = io.ReadFull(reader, body)
		if err != nil {
			fmt.Println("Error reading body:", err)
			return
		}
	}

	fmt.Println(body)
	fmt.Printf("Body (%d bytes): %s\n", bodyLength, string(body))

}

// получить нттр запрос, сохранить, прочитать заголовок и боди, вывести на экран, ответить
