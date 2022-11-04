package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	ip   = "127.0.0.1"
	port = "1234"
)

func main() {

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	address := fmt.Sprintf("%s:%s", ip, port)
	startNetService(address)

	<-sig
}

func startNetService(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("failed to listen")
	}
	fmt.Println("listening...")

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("failed to accept")
				continue
			}
			fmt.Println("connected")

			go func() {
				buf := make([]byte, 1024)
				for {
					l, err := conn.Read(buf)
					if err != nil || l == 0 {
						fmt.Println("failed to read")
						return
					}

					fmt.Println("from client:", string(buf))
					buf := []byte("hello I am server")
					_, err = conn.Write(buf)
					if err != nil {
						fmt.Println("failed to read")
					}
				}
			}()
		}
	}()
}
