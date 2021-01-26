package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":32412")
	if err != nil {
		panic(err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {

		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			for {
				message, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					log.Printf("Error: %+v", err.Error())
					return
				}

				log.Println("Message:", string(message))
				fmt.Println(string(message))
			}
		}(conn)
	}
}
