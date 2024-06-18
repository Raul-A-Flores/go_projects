package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	l, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal("couldnt listen to network port %v", l)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("err while accept", err)
		}

		go handle(conn)

	}
}

func handle(conn net.Conn) {
	fmt.Println("connected to : ", conn.RemoteAddr())
	for {
		var buffer [1024]byte
		_, err := conn.Read(buffer[:])
		if err != nil {
			//log.Fatalln("err while accept", err)
			log.Printf("err while readong from conn: %v, exiting....", err)
			return
		}
		fmt.Println("message read", string(buffer[:]))
	}
}
