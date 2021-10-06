package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

type clientInfo struct {
	messages client
	name     string
}

var (
	entering     = make(chan client)
	leaving      = make(chan client)
	ch_input_all = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	ch_client := make(chan string)

	go clientWriter(conn, ch_client)

	input := bufio.NewScanner(conn)

	client_ip := conn.RemoteAddr().String()

	entering <- ch_client

	log.Println(client_ip + " has arrived")
	b := true
	name := ""
	for input.Scan() {
		if b {
			b = false
			name = input.Text()
			fmt.Println("Hello clinet ", name)

			ch_client <- "Client from:" + client_ip + " name " + name
			//ch_input_all <- client_ip + " name " + name + " has arrived"
		} else {
			ch_input_all <- name + ": " + input.Text()
		}
	}

	leaving <- ch_client
	ch_input_all <- client_ip + " has left"
	conn.Close()
}

//Send responce to client
func clientWriter(conn net.Conn, ch_client <-chan string) {
	for msg := range ch_client {
		str := fmt.Sprintln(">>", msg)
		fmt.Fprintln(conn, str)
		fmt.Println(str)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-ch_input_all:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
