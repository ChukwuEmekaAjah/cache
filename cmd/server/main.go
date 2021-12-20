package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/ChukwuEmekaAjah/cache"
	"github.com/ChukwuEmekaAjah/cache/internal/parser"
)

type insertion struct {
	key    string
	values []string
}

var cacheMap = make(map[string]*parser.KeyValue)

// var values = map[string]insertion{}

func main() {
	cache.Config()
	arguments := os.Args

	portAddress := ":1996"

	if len(arguments) > 1 {
		portAddress = ":" + arguments[1]
	}

	l, err := net.Listen("tcp", portAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go tcpHandler(c)
	}
}

func tcpHandler(c net.Conn) {

	defer c.Close()
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.ToUpper(strings.TrimSpace(string(netData))) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		if strings.TrimSpace(netData) == "" {
			fmt.Println("Closing TCP server!")
			c.Close()
			return
		}

		isValidCommand := parser.Parse(strings.TrimSpace(netData))
		fmt.Println("command is ", string(netData), isValidCommand)
		if !isValidCommand {
			c.Write([]byte("Invalid command sent"))
		}

		commandParts := strings.Fields(strings.TrimSpace(netData))

		commandAction, exists := parser.ParserFunctions[strings.ToUpper(commandParts[0])]
		retrievalAction, ok := parser.RetrievalFunctions[strings.ToUpper(commandParts[0])]

		if exists == false && ok == false {
			c.Write([]byte("Invalid command sent 2"))
		}

		if exists {
			parsedValue := commandAction(strings.ToUpper(commandParts[0]), commandParts[1:], cacheMap)

			cacheMap[commandParts[1]] = parsedValue
		}

		// retrieval function
		if ok {
			parsedValue, err := retrievalAction(commandParts[0], commandParts[1:], cacheMap)

			if err != nil {
				c.Write([]byte("Invalid command sent 3"))
			}
			fmt.Println("retrieved value", parsedValue)
			c.Write([]byte(parsedValue))
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := "\n" + t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
