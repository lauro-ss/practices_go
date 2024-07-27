package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.5:7878")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	scannerServer := bufio.NewScanner(conn)
	go func(s *bufio.Scanner) {
		for s.Scan() {
			fmt.Println("mensagem:", scannerServer.Text())
		}
	}(scannerServer)

	for scanner.Scan() {
		if _, err := conn.Write([]byte(scanner.Text() + "\n")); err != nil {
			conn.Close()
			log.Fatal(err)
		}
	}
	conn.Close()
}
