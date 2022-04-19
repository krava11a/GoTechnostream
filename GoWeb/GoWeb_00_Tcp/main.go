package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":5000")

	for {
		//ждем пока не придет клиент
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Can not connect!!")
			conn.Close()
			continue
		}

		fmt.Println("Connected")

		//создаем Reader для чтения инфы из сокета
		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading!")

		go func (conn net.Conn){
			defer conn.Close()

			for true {
				rbyte, err := bufReader.ReadByte()

				if err != nil {
					fmt.Println("Can not read!", err)
					break
				}

				fmt.Print(string(rbyte))
				conn.Write([]byte("recieved"))
			}

		}(conn)

	}
}
