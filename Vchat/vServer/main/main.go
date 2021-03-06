package main

import (
	"Vchat/vServer/model"
	"fmt"
	"log"
	"net"
	"time"
)

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	// init pool
	initPool("localhost:6379", 8, 0, 100*time.Second)

	// init userDao
	initUserDao()

	//172.17.19.243
	listener, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		log.Fatalln("listen failed", err)
	}
	defer listener.Close()

	// accept connect on loop
	fmt.Println("welcome to Vchat")
	fmt.Println()
	fmt.Println("waiting for first connection")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server connected failed", err)
		}

		up := &Processor{
			Conn: conn,
		}

		// generate new goroutine when got a connection
		go up.Process()
	}
}
