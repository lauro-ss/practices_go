package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	network, err := net.Listen("tcp", "127.0.0.1:7878")
	if err != nil {
		log.Fatal(err)
	}
	defer network.Close()

	for {
		conn, err := network.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			reader := bufio.NewReader(c)
			for {
				m, err := reader.ReadString('\n')
				if err != nil {
					c.Close()
					return
				}
				fmt.Println(m)
				c.Write([]byte("conection"))
			}
		}(conn)
	}
}

func print() {
	fmt.Println("Hello World")
}
