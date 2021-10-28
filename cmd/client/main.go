package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args

	defaultServerAddress := "localhost:1996"

	if len(os.Args) > 1 {
		defaultServerAddress = arguments[1]
	}

	c, err := net.Dial("tcp", defaultServerAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		for {
			message, err := bufio.NewReader(c).ReadString('\n')

			if err != nil {
				break
			}
			fmt.Print("->: " + message)
		}

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
