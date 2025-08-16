package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"
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
	file, err := os.Create("bodylog.txt")

	if err != nil {
		fmt.Println("Ошибка при создании файла")
	}
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
	file.Write(header)
	fmt.Println(header)

	// contentLengh, err := strconv.Atoi(requestline)
	// if err != nil {
	// 	fmt.Print(" Ошибка Атои   ")
	// }

}

// получить нттр запрос, сохранить, прочитать заголовок и боди, вывести на экран, ответить
