package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ip   = "127.0.0.1"
	port = "1234"
)

func main() {

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	address := fmt.Sprintf("%s:%s", ip, port)
	for i := 0; i < 1000; i++ {
		go startNetService(address)
	}

	<-sig
}

func startNetService(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("failed to connect")
		return
	}

	fmt.Println("connected")

	go func() {
		buf := make([]byte, 1024)
		for {
			l, err := conn.Read(buf)
			if err != nil || l == 0 {
				fmt.Println("failed to read")
				err = conn.Close()
				return
			}

			fmt.Println("from server:", string(buf))
		}
	}()

	for {
		time.Sleep(time.Second)

		buf := []byte("hello I am client")
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println("failed to write")
			err = conn.Close()
		}
	}

}
