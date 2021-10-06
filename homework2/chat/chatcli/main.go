package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Your name is: ", name)

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	conn.Write([]byte(name))

	go func() {
		//os.Stdout.WriteString("hello " + name + "\n")
		io.Copy(os.Stdout, conn)
	}()

	io.Copy(conn, os.Stdin) // until you send ^Z
	fmt.Printf("%s: exit", conn.LocalAddr())
}
