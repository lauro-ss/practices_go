package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	network, err := net.Listen("tcp", "192.168.0.5:7878")
	if err != nil {
		log.Fatal(err)
	}
	defer network.Close()
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	//chats []string

	for {
		conn, err := network.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			fmt.Println(c.LocalAddr(), c.RemoteAddr())
			scan := bufio.NewScanner(c)
			scanClient := bufio.NewScanner(os.Stdin)
			go func(s *bufio.Scanner, con net.Conn) {
				for scanClient.Scan() {
					con.Write([]byte(scanClient.Text() + "\n"))
				}
			}(scanClient, c)
			for scan.Scan() {
				fmt.Println("mensagem:", scan.Text())
			}
			conn.Close()
		}(conn)
	}
}

func print() {
	fmt.Println("Hello World")
}
